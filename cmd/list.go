package cmd

import (
	"os"
	"text/tabwriter"

	"github.com/nerfthisdev/todolite/internal/storagemodule"
	"github.com/spf13/cobra"
)

var all bool

var listId bool

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Shows a list of current tasks",
	Long:  `Lists your to-do tasks`,
	Run: func(cmd *cobra.Command, args []string) {
		if all {
			listAll()
		} else {
			listNotCompleted()
		}
	},
}

func listAll() {

	tasks, _ := storagemodule.LoadStorage()
	w := tabwriter.NewWriter(os.Stdout, 0, 4, 8, ' ', 0)
	printHeader(w)
	for _, task := range tasks {
		printTaskRow(w, task)
	}
	w.Flush()
}

func listNotCompleted() {
	tasks, _ := storagemodule.LoadStorage()
	w := tabwriter.NewWriter(os.Stdout, 0, 4, 6, ' ', 0)
	printHeader(w)
	for _, task := range tasks {
		if !task.Status {
			printTaskRow(w, task)
		}
	}
	w.Flush()

}

func printHeader(w *tabwriter.Writer) {
	if listId {
		w.Write([]byte("Id\tStatus\tDescription\tDateCreated\n"))
	} else {
		w.Write([]byte("Status\tDescription\tDateCreated\n"))
	}
}

func printTaskRow(w *tabwriter.Writer, task storagemodule.Task) {
	if listId {
		w.Write([]byte(task.StringWithId() + "\n"))
	} else {
		w.Write([]byte(task.String() + "\n"))
	}
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&all, "all", "a", false, "Include all items")
	listCmd.Flags().BoolVarP(&listId, "id", "i", false, "Show id all items")
}
