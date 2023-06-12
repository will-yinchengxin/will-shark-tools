package ws

import "github.com/spf13/cobra"

var CreateServiceCmd = &cobra.Command{
	Use:     "service",
	Short:   "create a new controller",
	Example: "WS c service user",
	Args:    cobra.ExactArgs(1),
	Run:     runCreate,
}
