/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/nerfthisdev/todolite/internal/storagemodule"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task to your todo list!",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		tasks, _ := storagemodule.LoadStorage()
		id := tasks[len(tasks)-1].Id + 1
		fmt.Print("Name of your task: ")
		reader := bufio.NewReader(os.Stdin)
		name, _ := reader.ReadString('\n')
		name = strings.TrimSuffix(name, "\n")

		storagemodule.AppendTaskToCSV(storagemodule.Task{
			Id:           id,
			Name:         name,
			Status:       false,
			Creationdate: time.Now(),
		})

	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
