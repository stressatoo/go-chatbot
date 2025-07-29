package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var askCmd = &cobra.Command{
	Use:     "ask",
	Aliases: []string{"chat"},
	Short:   "Talk to Agent",
	Long:    "Ask a question to your agent",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("User: %s\n", args[0])
		fmt.Printf("iGor: ")
		fmt.Printf("%s", Ask(args[0]))
	},
}

func init() {
	rootCmd.AddCommand(askCmd)
}
