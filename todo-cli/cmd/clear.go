package cmd

import (
	"fmt"

	"todo-cli/models"
	"todo-cli/storage"

	"github.com/spf13/cobra"
)

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		// Overwrite storage with an empty slice
		if err := storage.SaveTasks([]models.Task{}); err != nil {
			fmt.Println("Error clearing tasks:", err)
			return
		}
		fmt.Println("All tasks cleared.")
	},
}

func init() {
	rootCmd.AddCommand(clearCmd)
}
