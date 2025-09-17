package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"todo-cli/storage" // only if you need the type

	"github.com/spf13/cobra"
)

var printCmd = &cobra.Command{
	Use:   "print",
	Short: "Prints task",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := storage.LoadTasks()
		if err != nil {
			fmt.Println("Error loading tasks:", err)
			return
		}
		// create a tab writer that writes to stdout
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

		// header
		fmt.Fprintln(w, "ID\tTitle\tCompleted\tCreated\tAge")

		// rows
		for _, t := range tasks {
			age := time.Since(t.CreatedAt).Round(time.Minute) // human-friendly
			fmt.Fprintf(w, "%d\t%s\t%t\t%s\t%s\n",
				t.ID, t.Title, t.Completed, t.CreatedAt.Format("2006-01-02 15:04"), age)
		}

		// flush to output
		w.Flush()

	},
}

func init() {
	rootCmd.AddCommand(printCmd)
}
