// Code generated by goa v3.2.4, DO NOT EDIT.
//
// entity_hall HTTP server
//
// Command:
// $ goa gen boot/design

package server

import (
	entityhall "boot/gen/entity_hall"
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Server lists the entity_hall service endpoint HTTP handlers.
type Server struct {
	Mounts           []*MountPoint
	WaitLineOverview http.Handler
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the design.
type ErrorNamer interface {
	ErrorName() string
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the entity_hall service endpoints
// using the provided encoder and decoder. The handlers are mounted on the
// given mux using the HTTP verb and path defined in the design. errhandler is
// called whenever a response fails to be encoded. formatter is used to format
// errors returned by the service methods prior to encoding. Both errhandler
// and formatter are optional and can be nil.
func New(
	e *entityhall.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"WaitLineOverview", "POST", "/api/entity_hall/get_wait_line_overview"},
		},
		WaitLineOverview: NewWaitLineOverviewHandler(e.WaitLineOverview, mux, decoder, encoder, errhandler, formatter),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "entity_hall" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.WaitLineOverview = m(s.WaitLineOverview)
}

// Mount configures the mux to serve the entity_hall endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountWaitLineOverviewHandler(mux, h.WaitLineOverview)
}

// MountWaitLineOverviewHandler configures the mux to serve the "entity_hall"
// service "WaitLineOverview" endpoint.
func MountWaitLineOverviewHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/api/entity_hall/get_wait_line_overview", f)
}

// NewWaitLineOverviewHandler creates a HTTP handler which loads the HTTP
// request and calls the "entity_hall" service "WaitLineOverview" endpoint.
func NewWaitLineOverviewHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeWaitLineOverviewRequest(mux, decoder)
		encodeResponse = EncodeWaitLineOverviewResponse(encoder)
		encodeError    = EncodeWaitLineOverviewError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "WaitLineOverview")
		ctx = context.WithValue(ctx, goa.ServiceKey, "entity_hall")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}
