/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// lstDptCmd représente la commande lstDpt
var lstdCmd = &cobra.Command{
	Use:   "lstd",
	Short: "Cette commande permet de lister tous les départements existants dans le fichier Source",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Recherche des départements dans le fichier source...")
		// Appel de la fonction pour lister les départements
		my_WI.ExtractDepartFromCSV()
	},
}

// Fonction lors de l'appel de la commande
func ExecuteLstDpt() error {
	return lstdCmd.Execute()
}

func init() {
	RootCmd.AddCommand(lstdCmd)
}
