/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/abdealt/meliodas/components"
	"github.com/spf13/cobra"
)

// statCmd represents the stat command
var statCmd = &cobra.Command{
	Use:   "stat",
	Short: "Cette commande sert a afficher les statistique obtenus lors du dernier traitement.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Commande 'stat' est écécuté\n")
		// TODO: Implement your logic here
		fmt.Printf("Il y'a au total %v éléments dans le fichier source. %v éléments de ce fichier ont été extraits.", components.ComptTotal, components.ComptElement)
	},
}

func init() {
	rootCmd.AddCommand(statCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// statCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// statCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
