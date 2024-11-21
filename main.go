/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"log"

	"github.com/nerfthisdev/todolite/internal/csvmodule"
)

func main() {
	// cmd.Execute()
	t, err := csvmodule.LoadStorage("zavupa.csv")

	if err != nil {
		log.Fatal("Error getting tasks")
	}

	for _, task := range t {
		fmt.Println(task.Id, task.Name, task.Status, task.Creationdate)
	}

	fmt.Println()
}
