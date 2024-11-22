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

		fmt.Print("Name of your task: ")
		reader := bufio.NewReader(os.Stdin)
		name, _ := reader.ReadString('\n')
		name = strings.TrimSuffix(name, "\n")

		storagemodule.AppendTaskToDB(storagemodule.Task{
			Name:         name,
			Status:       false,
			Creationdate: time.Now(),
		})

	},
}

func init() {
	rootCmd.AddCommand(addCmd)

}
