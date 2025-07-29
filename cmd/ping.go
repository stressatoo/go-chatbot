package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var pingCmd = &cobra.Command{
	Use:     "ping",
	Short:   "Talk to Agent",
	Long:    "Ask a question to your agent",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("PONG")

	},
}

func init() {
	rootCmd.AddCommand(pingCmd)
}
