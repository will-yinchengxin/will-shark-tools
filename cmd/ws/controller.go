package ws

import "github.com/spf13/cobra"

var CreateControllerCmd = &cobra.Command{
	Use:     "controller",
	Short:   "create a new controller",
	Example: "WS c controller user",
	Args:    cobra.ExactArgs(1),
	Run:     runCreate,
}
