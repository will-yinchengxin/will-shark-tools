package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"ws/cmd/ws"
)

const Version = "1.0.0"

var rootCmd = &cobra.Command{
	Use:     "c",
	Example: "WS c all 「user」",
	Version: "@copyright Copyright (C) 2023 Released under the MIT License\nauthor：Will Yin [https://github.com/will-yinchengxin]\nVersion: " + Version,
}

var createCmd = &cobra.Command{
	Use:     "c [type] [handler-name]",
	Example: "WS c controller user",
	Short:   "create a new controller/service/dao/entity for will-shark",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.AddCommand(ws.CreateAllCmd)
	createCmd.AddCommand(ws.CreateControllerCmd)
	createCmd.AddCommand(ws.CreateServiceCmd)
	createCmd.AddCommand(ws.CreateRouterCmd)
	createCmd.AddCommand(ws.CreateDaoCmd)
	createCmd.AddCommand(ws.CreateEntityCmd)
	createCmd.AddCommand(ws.CreateRequestCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
