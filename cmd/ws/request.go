package ws

import "github.com/spf13/cobra"

var CreateRequestCmd = &cobra.Command{
	Use:     "req",
	Short:   "create a new request",
	Example: "WS c req user",
	Args:    cobra.ExactArgs(1),
	Run:     runCreate,
}
