package main

//func TimerCmd() *cobra.Command {
//	serverCmd := &cobra.Command{
//		Use:   "runtimer",
//		Short: "启动定时更行中间库 up_pro_accept 表的 TimeLimit值 的定时任务",
//		Long:  "启动定时更行中间库 up_pro_accept 表的 TimeLimit值 的定时任务",
//		RunE: func(cmd *cobra.Command, args []string) error {
//			tool := tools.NewTimerTool()
//			err := tool.UpdateProAcceptTimeLimit(config.C)
//			if err != nil {
//				log.Fatalln(err)
//				return err
//			}
//			return nil
//		},
//		PreRunE: func(cmd *cobra.Command, args []string) error {
//			config.Init(cfgFile)
//			dao.InitDB(config.C)
//			return nil
//		},
//	}
//
//	serverCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")
//	return serverCmd
//}
