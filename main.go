/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"log"

	"github.com/leigme/gpf/cmd"
)

func init() {
	log.SetFlags(log.Llongfile | log.LstdFlags)
}

func main() {
	cmd.Execute()
}
