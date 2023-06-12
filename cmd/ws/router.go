package ws

import "github.com/spf13/cobra"

var CreateRouterCmd = &cobra.Command{
	Use:     "router",
	Short:   "create a new router",
	Example: "WS c router user",
	Args:    cobra.ExactArgs(1),
	Run:     runCreate,
}
