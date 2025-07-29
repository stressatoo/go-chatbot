package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "os"
)

var rootCmd = &cobra.Command{
    Use:   "igor",
    Short: "igor is a cli tool for creating and chatting with simple personalized chatbot",
    Long:  "igor is a cli tool for creating and chatting with simple personalized chatbot. Create and conversate with them!",
    Run: func(cmd *cobra.Command, args []string) {

    },
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintf(os.Stderr, "Oops. An error while executing iGor '%s'\n", err)
        os.Exit(1)
    }
}