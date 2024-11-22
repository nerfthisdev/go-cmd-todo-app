package cmd

import (
	"log"
	"os"
	"path/filepath"

	"github.com/nerfthisdev/todolite/internal/storagemodule"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "todolite",
	Short: "A to-do app right in your terminal!",
	Long:  `Todolite is a terminal app that lets you manage your to do list using simple commands`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("failed to get home directory: %v", err)
	}

	dbPath := filepath.Join(homeDir, ".todolite.db")
	err = storagemodule.InitDB(dbPath)
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}
}
