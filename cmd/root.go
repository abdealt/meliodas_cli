package cmd

import (
	"fmt"
	"os"

	"github.com/abdealt/meliodas_cli/config"
	"github.com/common-nighthawk/go-figure"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   config.AppName,
	Short: config.AppDisplayName,
	Long:  config.AppDesc,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	myFigure := figure.NewFigure(config.AppDisplayName, "", true)
	myFigure.Print()
}
