package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd is the base command
var rootCmd = &cobra.Command{
	Use:   "todo-cli",
	Short: "A simple CLI to manage your tasks",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to todo-cli! Use -h for help.")
	},
}

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
