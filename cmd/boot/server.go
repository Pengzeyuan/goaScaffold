package starter

import (
	"context"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"net/url"
	"os"
	"os/signal"
	"sync"

	"boot/config"
	controller "boot/controller"
	"boot/gen/log"
	"boot/gen/user"

	metricsMlwr "git.chinaopen.ai/yottacloud/go-libs/goa-libs/middleware/metrics"
	"git.chinaopen.ai/yottacloud/go-libs/panichandler"
	"go.uber.org/zap"

	entityhall "boot/gen/entity_hall"

	actualtime "boot/gen/actual_time"

	thirdpart "boot/gen/third_part"

	importfile "boot/gen/import_file"

	simulation "boot/gen/simulation"
)

func RunServer(cfg *config.Config, metrics *metricsMlwr.Prometheus) {

	// Setup logger. Replace logger with your own log package of choice.
	var (
		logger *log.Logger
	)
	{
		logger = log.New("starter", !cfg.Debug)
	}

	// Initialize the services.
	var (
		userSvc       user.Service
		entityHallSvc entityhall.Service
		actualTimeSvc actualtime.Service
		thirdpartSvc  thirdpart.Service
		importFileSvc importfile.Service
		simulationSvc simulation.Service
	)
	{
		userSvc = controller.NewUser(logger)
		entityHallSvc = controller.NewEntityHall(logger)
		actualTimeSvc = controller.NewActualTime(logger)
		thirdpartSvc = controller.NewThirdPart(logger)
		importFileSvc = controller.NewImportFile(logger)
		simulationSvc = controller.NewSimulation(logger)
	}

	// Wrap the services in endpoints that can be invoked from other services
	// potentially running in different processes.
	var (
		userEndpoints       *user.Endpoints
		entityHallEndpoints *entityhall.Endpoints
		actualTimeEndpoints *actualtime.Endpoints
		thirdPartEndpoints  *thirdpart.Endpoints
		importFileEndpoints *importfile.Endpoints
		simulationEndpoints *simulation.Endpoints
	)
	{
		userEndpoints = user.NewEndpoints(userSvc)
		entityHallEndpoints = entityhall.NewEndpoints(entityHallSvc)
		actualTimeEndpoints = actualtime.NewEndpoints(actualTimeSvc)
		thirdPartEndpoints = thirdpart.NewEndpoints(thirdpartSvc)
		importFileEndpoints = importfile.NewEndpoints(importFileSvc)
		simulationEndpoints = simulation.NewEndpoints(simulationSvc)
	}

	// Create channel used by both the signal handler and server goroutines
	// to notify the main goroutine when to stop the server.
	errc := make(chan error)

	// Setup interrupt handler. This optional step configures the process so
	// that SIGINT and SIGTERM signals cause the services to stop gracefully.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		errc <- fmt.Errorf("%s", <-c)
	}()

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	addr := fmt.Sprintf("http://%s:%s", cfg.Server.Host, cfg.Server.HTTPPort)
	u, _ := url.Parse(addr)
	handleHTTPServer(ctx, u.Host, userEndpoints,
		entityHallEndpoints,
		actualTimeEndpoints,
		thirdPartEndpoints,
		importFileEndpoints,
		simulationEndpoints,
		&wg, errc, logger, metrics, cfg.Debug)

	// Wait for signal.
	logger.Infof("exiting (%v)", <-errc)

	// Send cancellation signal to the goroutines.
	cancel()

	wg.Wait()
	logger.Info("exited")
}

// 开启 pprof
func RunDebugPprofServer(addr string) {
	defer panichandler.ZapHandler(zap.L()).Handle()
	zap.L().Sugar().Infof("启动 pprof 监听 %s.", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		zap.L().Error("开启 pprof 监听失败 %s", zap.Error(err))
	}
}
