package ws

import "github.com/spf13/cobra"

var CreateAllCmd = &cobra.Command{
	Use:     "all",
	Short:   "create all",
	Example: "WS c all user",
	Args:    cobra.ExactArgs(1),
	Run:     runCreate,
}
