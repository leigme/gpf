/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/leigme/gpf/config"
	"github.com/leigme/gpf/model"
	"github.com/spf13/cobra"
)

var (
	p  = model.Param{}
	cj = config.Json{}
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gpf",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		defer cj.Update()
		bindLast()
		if err := paramCheck(); err != nil {
			log.Fatalln(err)
		}
		generate()
		cj.LastTemplate = p.Template
		cj.LastGenerate = p.Generate
	},
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
	createDir()
	cj.Load()
	rootCmd.PersistentFlags().StringVar(&p.Template, "t", cj.LastTemplate, "")
	rootCmd.PersistentFlags().StringVar(&p.Args, "a", "", "")
	rootCmd.PersistentFlags().StringVar(&p.Generate, "g", cj.LastGenerate, "")
}

func createDir() {
	configPath := config.Path(".config/gpf", "conf.json")
	_, err := os.Stat(filepath.Dir(configPath))
	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(configPath, os.ModePerm)
			if err == nil {
				return
			}
		}
		log.Fatalln(err)
	}
}

func bindLast() {
	if strings.EqualFold(p.Template, "") && !strings.EqualFold(cj.LastTemplate, "") {
		p.Template = cj.LastTemplate
	}
	if strings.EqualFold(p.Generate, "") && !strings.EqualFold(cj.LastGenerate, "") {
		p.Generate = cj.LastGenerate
	}
}

func paramCheck() error {
	if strings.EqualFold(p.Template, "") {
		return errors.New("--t is nil")
	}
	if strings.EqualFold(p.Generate, "") {
		return errors.New("--g is nil")
	}

	if strings.EqualFold(p.Args, "") {
		return errors.New("--a is nil")
	}
	return nil
}

func generate() {
	log.Printf("param: %s", p)
}
