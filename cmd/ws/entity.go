package ws

import "github.com/spf13/cobra"

var CreateEntityCmd = &cobra.Command{
	Use:     "entity",
	Short:   "create a new entity",
	Example: "WS c entity user",
	Args:    cobra.ExactArgs(1),
	Run:     runCreate,
}
