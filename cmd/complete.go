package cmd

import (
	"log"
	"strconv"

	"github.com/nerfthisdev/todolite/internal/storagemodule"
	"github.com/spf13/cobra"
)

var idFlag bool

var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Mark task as completed!",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if idFlag {
			if len(args) == 0 {
				log.Fatal("No arguments found")
			}
			id, err := strconv.Atoi(args[0])

			if err != nil {
				log.Fatal("Couldn't parse id argument")
			}

			storagemodule.UpdateTaskStatusById(id, true)
		} else {
			if len(args) == 0 {
				log.Fatal("No arguments found")
			}

			if err := storagemodule.UpdateTaskStatusByName(args[0], true); err != nil {
				log.Fatal("Couldnt update tasks")
			}

		}
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
	completeCmd.Flags().BoolVarP(&idFlag, "id", "i", false, "Use Id to complete a task")
}
