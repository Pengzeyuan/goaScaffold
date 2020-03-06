package main

import (
	"log"

	"git.chinaopen.ai/iot/go-libs/redis"
	"github.com/lneoe/go-help-libs/version"
	"github.com/spf13/cobra"

	starterCmd "starter/cmd/starter"
	"starter/config"
	"starter/dao"

	metricsMlwr "starter/middleware/metrics"
)

var (
	cfgFile string
)

func serverCmd() *cobra.Command {
	serverCmd := &cobra.Command{
		Use: "runserver",
		RunE: func(cmd *cobra.Command, args []string) error {
			var metrics *metricsMlwr.Prometheus
			if config.C.Metrics.Enabled {
				metrics = metricsMlwr.NewPrometheus("starter", nil)
				go metrics.Start(config.C.Metrics.Addr)
			}

			// pprof 这里启动失败会直接 panic
			if config.C.Pprof.Enabled {
				go starterCmd.RunDebugPprofServer(config.C.Pprof.Addr)
			}

			starterCmd.RunServer(config.C, metrics)

			return nil
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			config.Init(cfgFile)
			dao.InitDB(config.C)
			dao.AutoMigrateDB()
			if err := redis.Connect(); err != nil {
				return err
			}

			return nil
		},
	}

	serverCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")

	return serverCmd
}

func versionCmd() *cobra.Command {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "show version",
		Long:  "show version",
		RunE: func(cmd *cobra.Command, args []string) error {
			version.PrintVersion()
			return nil
		},
	}

	return versionCmd
}

func main() {
	var RootCmd = cobra.Command{
		Use: "starter",
	}

	RootCmd.AddCommand(serverCmd())
	RootCmd.AddCommand(versionCmd())

	if err := RootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
