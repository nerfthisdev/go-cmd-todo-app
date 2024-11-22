/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"text/tabwriter"

	"github.com/nerfthisdev/todolite/internal/storagemodule"
	"github.com/spf13/cobra"
)

var all bool

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
	w.Write([]byte("Id\tName\tStatus\tDateCreated\n"))
	for _, task := range tasks {
		w.Write([]byte(task.String() + "\n"))

	}
	w.Flush()
}

func listNotCompleted() {
	tasks, _ := storagemodule.LoadStorage()
	w := tabwriter.NewWriter(os.Stdout, 0, 4, 8, ' ', 0)
	w.Write([]byte("Id\tName\tStatus\tDateCreated\n"))
	for _, task := range tasks {
		if !task.Status {
			w.Write([]byte(task.String() + "\n"))
		}
	}
	w.Flush()

}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&all, "all", "a", false, "Include all items")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
