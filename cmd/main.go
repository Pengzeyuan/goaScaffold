package main

import (
	"boot/utils"
	"log"

	"github.com/lneoe/go-help-libs/version"
	"github.com/spf13/cobra"

	bootCmd "boot/cmd/boot"
	"boot/config"
	"boot/dao"

	metricsMlwr "git.chinaopen.ai/yottacloud/go-libs/goa-libs/middleware/metrics"
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
				go bootCmd.RunDebugPprofServer(config.C.Pprof.Addr)
			}

			bootCmd.RunServer(config.C, metrics)

			return nil
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			config.Init(cfgFile)
			dao.InitDB(config.C)
			//dao.InitCanal(config.C)
			dao.AutoMigrateDB()
			if err := utils.ConnectNats(); err != nil {
				log.Fatalln(err)
				return err
			}
			//if err := redis.Connect(); err != nil {
			//	return err
			//}

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

var RootCmd cobra.Command = cobra.Command{
	Use: "Scaffold",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func main() {

	RootCmd.AddCommand(serverCmd())
	RootCmd.AddCommand(versionCmd())
	RootCmd.AddCommand(userCmd())
	RootCmd.AddCommand(WebsocketCmd())
	RootCmd.AddCommand(CanalCmd())
	if err := RootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
