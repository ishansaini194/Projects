package cmd

import (
	"fmt"

	"time"
	"todo-cli/models"
	"todo-cli/storage"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create [task]",
	Short: "Create a new task",
	Args:  cobra.ExactArgs(1), // ensures exactly 1 argument is passed
	Run: func(cmd *cobra.Command, args []string) {
		task := args[0]
		fmt.Println("Created task:", task)
		tasks, _ := storage.LoadTasks()
		newTask := models.Task{
			ID:        len(tasks) + 1,
			Title:     args[0],
			Completed: false,
			CreatedAt: time.Now(),
		}
		tasks = append(tasks, newTask)
		storage.SaveTasks(tasks)

	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
