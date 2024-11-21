/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/nerfthisdev/todolite/internal/functions"
	"github.com/spf13/cobra"
)

// timefromzoneCmd represents the timefromzone command
var timefromzoneCmd = &cobra.Command{
	Use:   "timefromzone",
	Short: "Tells a time when given a location",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		timefromzone := args[0]
		timeNow, err := functions.GetTimeFromZone(timefromzone)
		if err != nil {
			fmt.Println("Неверная зона")
		}

		fmt.Println(timeNow)
	},
}

func init() {
	rootCmd.AddCommand(timefromzoneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// timefromzoneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// timefromzoneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
