package starter

import (
	controller "boot/controller"
	"context"
	"net/http"
	"os"
	"sync"
	"time"

	"git.chinaopen.ai/yottacloud/go-libs/goa-libs/middleware/metrics"
	"go.uber.org/zap"
	goahttp "goa.design/goa/v3/http"
	httpmdlwr "goa.design/goa/v3/http/middleware"
	"goa.design/goa/v3/middleware"

	"boot/gen/log"
	mdlwr "boot/middleware"

	usersvr "boot/gen/http/user/server"
	"boot/gen/user"

	entityhall "boot/gen/entity_hall"
	entityhallsvr "boot/gen/http/entity_hall/server"

	actualtime "boot/gen/actual_time"
	actualtimesvr "boot/gen/http/actual_time/server"

	thirdpartsvr "boot/gen/http/third_part/server"
	thirdpart "boot/gen/third_part"

	importfilesvr "boot/gen/http/import_file/server"
	importfile "boot/gen/import_file"

	simulationsvr "boot/gen/http/simulation/server"
	simulation "boot/gen/simulation"
)

// handleHTTPServer starts configures and starts a HTTP server on the given
// URL. It shuts down the server if any error is received in the error channel.
func handleHTTPServer(ctx context.Context, host string,
	userEndpoints *user.Endpoints,
	entityhallEndpoints *entityhall.Endpoints,
	actualTimeEndpoints *actualtime.Endpoints,
	thirdpartEndpoints *thirdpart.Endpoints,
	importFileEndpoints *importfile.Endpoints,
	simulationEndpoints *simulation.Endpoints,
	wg *sync.WaitGroup, errc chan error,
	logger *log.Logger, metrics *metrics.Prometheus, debug bool) {

	// Setup goa log adapter.
	var (
		adapter middleware.Logger
	)
	{
		adapter = logger
	}

	// Provide the transport specific request decoder and response encoder.
	// The goa http package has built-in support for JSON, XML and gob.
	// Other encodings can be used by providing the corresponding functions,
	// see goa.design/implement/encoding.
	var (
		dec = goahttp.RequestDecoder
		enc = goahttp.ResponseEncoder
	)

	// Build the service HTTP request multiplexer and configure it to serve
	// HTTP requests to the service endpoints.
	var mux goahttp.Muxer
	{
		mux = goahttp.NewMuxer()
	}

	// Wrap the endpoints with the transport specific layers. The generated
	// server packages contains code generated from the design which maps
	// the service input and output data structures to HTTP requests and
	// responses.
	var (
		userServer       *usersvr.Server
		entityhallServer *entityhallsvr.Server
		actualTimeServer *actualtimesvr.Server
		thirdpartServer  *thirdpartsvr.Server
		importFileServer *importfilesvr.Server
		simulationServer *simulationsvr.Server
	)
	{
		eh := errorHandler(logger)
		userServer = usersvr.New(userEndpoints, mux, dec, enc, eh, mdlwr.GoaErrorFormatterFunc)
		entityhallServer = entityhallsvr.New(entityhallEndpoints, mux, dec, enc, eh, mdlwr.GoaErrorFormatterFunc)
		actualTimeServer = actualtimesvr.New(actualTimeEndpoints, mux, dec, enc, eh, mdlwr.GoaErrorFormatterFunc)
		thirdpartServer = thirdpartsvr.New(thirdpartEndpoints, mux, dec, enc, eh, mdlwr.GoaErrorFormatterFunc)

		importFileServer = importfilesvr.New(importFileEndpoints, mux, dec, enc, eh, mdlwr.GoaErrorFormatterFunc, controller.FileImportDecoderFunc)
		simulationServer = simulationsvr.New(simulationEndpoints, mux, dec, enc, eh, mdlwr.GoaErrorFormatterFunc)
	}
	// Configure the mux.
	usersvr.Mount(mux, userServer)
	entityhallsvr.Mount(mux, entityhallServer)
	actualtimesvr.Mount(mux, actualTimeServer)
	thirdpartsvr.Mount(mux, thirdpartServer)
	importfilesvr.Mount(mux, importFileServer)
	simulationsvr.Mount(mux, simulationServer)
	// Wrap the multiplexer with additional middlewares. Middlewares mounted
	// here apply to all the service endpoints.
	var handler http.Handler = mux
	{
		if debug {
			handler = httpmdlwr.Debug(mux, os.Stdout)(handler)
		}

		handler = mdlwr.PopulateRequestContext()(handler)
		handler = httpmdlwr.RequestID()(handler)

		if metrics != nil {
			handler = metrics.HandlerFunc(adapter)(handler)
		} else {
			handler = httpmdlwr.Log(adapter)(handler)
		}
	}

	// Start HTTP server using default configuration, change the code to
	// configure the server as required by your service.
	srv := &http.Server{Addr: host, Handler: handler}
	for _, m := range userServer.Mounts {
		logger.Infof("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}

	for _, m := range entityhallServer.Mounts {
		logger.Infof("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}
	for _, m := range actualTimeServer.Mounts {
		logger.Infof("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}
	(*wg).Add(1)
	go func() {
		defer (*wg).Done()

		// Start HTTP server in a separate goroutine.
		go func() {
			logger.Infof("HTTP server listening on %q", host)
			errc <- srv.ListenAndServe()
		}()

		<-ctx.Done()
		logger.Infof("shutting down HTTP server at %q", host)

		// Shutdown gracefully with a 30s timeout.
		ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
		defer cancel()

		_ = srv.Shutdown(ctx)
	}()
}

// errorHandler returns a function that writes and logs the given error.
// The function also writes and logs the error unique ID so that it's possible
// to correlate.
func errorHandler(logger *log.Logger) func(context.Context, http.ResponseWriter, error) {
	return func(ctx context.Context, w http.ResponseWriter, err error) {
		id := ctx.Value(middleware.RequestIDKey).(string)
		_, _ = w.Write([]byte("[" + id + "] encoding: " + err.Error()))
		logger.With(zap.String("id", id)).Error(err.Error())
	}
}
