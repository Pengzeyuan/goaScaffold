package main

import (
	"boot/cmd/tools"
	"boot/config"
	"boot/dao"
	"github.com/spf13/cobra"
	"log"
)

// userCmd represents the user command
func userCmd() *cobra.Command {
	var (
		userName, password string
	)

	cmd := &cobra.Command{
		Use:   "create-user",
		Short: "创建用户",
		Long:  "创建用户",
		RunE: func(cmd *cobra.Command, args []string) error {
			tool := tools.AdmTool{}
			// 创建用户
			if err := tool.CreateAdm(userName, password); err != nil {
				log.Fatalln(err)
				return err
			}
			return nil
		},

		PreRunE: func(cmd *cobra.Command, args []string) error {
			config.Init(cfgFile)
			dao.InitDB(config.C)
			return nil
		},
	}
	cmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")
	// 输入的选项
	cmd.Flags().StringVarP(&userName, "userName", "u", "", "用户名")
	cmd.Flags().StringVarP(&password, "password", "p", "", "密码")

	_ = cmd.MarkFlagRequired("userName")
	_ = cmd.MarkFlagRequired("password")

	return cmd
}
