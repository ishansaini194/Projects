package cmd

import (
	"fmt"
	"strconv"

	"todo-cli/storage"

	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "complete [id]",
	Short: "Mark a task as completed",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		idStr := args[0]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Printf("Invalid id: %s\n", idStr)
			return
		}

		tasks, err := storage.LoadTasks()
		if err != nil {
			fmt.Println("Error loading tasks:", err)
			return
		}

		found := false
		for i := range tasks {
			if tasks[i].ID == id {
				found = true
				if tasks[i].Completed {
					fmt.Printf("Task %d is already completed.\n", id)
				} else {
					tasks[i].Completed = true
					if err := storage.SaveTasks(tasks); err != nil {
						fmt.Println("Error saving tasks:", err)
						return
					}
					fmt.Printf("Task %d marked completed.\n", id)
				}
				break
			}
		}

		if !found {
			fmt.Printf("Task %d not found.\n", id)
		}
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
