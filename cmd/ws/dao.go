package ws

import "github.com/spf13/cobra"

var CreateDaoCmd = &cobra.Command{
	Use:     "dao",
	Short:   "create a new dao",
	Example: "WS c dao user",
	Args:    cobra.ExactArgs(1),
	Run:     runCreate,
}
