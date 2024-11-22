/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/nerfthisdev/todolite/cmd"
	"github.com/nerfthisdev/todolite/internal/storagemodule"
)

func main() {
	cmd.Execute()
	defer storagemodule.DB.Close()
}
