package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/abdealt/meliodas_cli/config"
	"github.com/common-nighthawk/go-figure"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "meliodas",
	Short: "Meliodas CLI",
	Long:  `Meliodas CLI est une application permettant d'exécuter des commandes pour extraire des données depuis un fichier CSV.`,
	// Exécution avant toute commande pour charger les variables d'environnement
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Charger le fichier .env qui est dans le même répertoire que root.go
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Erreur lors du chargement du fichier .env : %v", err)
		}
		fmt.Println("Fichier .env chargé avec succès")
	},
}

// Exécuter rootCmd
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Afficher le logo Meliodas
	myFigure := figure.NewFigure(config.AppDisplayName, "", true)
	myFigure.Print()
}
