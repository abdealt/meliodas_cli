package cmd

import (
	"log"
	"os"

	"github.com/abdealt/meliodas_cli/config"
	"github.com/common-nighthawk/go-figure"
	"github.com/spf13/cobra"
)

// Variables globales
var TotalElements int
var ExtractedElements int

// Définir la commande root
var rootCmd = &cobra.Command{
	Use:   "meliodas",
	Short: "\nUn CLI pour lire un fichier",
	Long:  "\nMeliodas CLI est un outil en ligne de commande simple et efficace conçu pour automatiser le traitement de fichiers CSV.\nIl permet aux utilisateurs d'extraire rapidement des données spécifiques à partir de fichiers, sans avoir à ouvrir un logiciel complexe.",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Chemin du fichier .env (à ajuster selon le fichier)
		filePath := "C:\\Users\\Utilisateur\\Desktop\\MELIODAS\\meliodas_cli\\.env"

		// Ouvrir le fichier .env
		file, err := os.Open(filePath)
		if err != nil {
			log.Fatalf("\nErreur lors de l'ouverture du fichier : %v", err)
		}
		defer file.Close()
	},
}

func init() {
	myLogo := figure.NewFigure(config.AppDisplayName, "", true)
	myLogo.Print()
}

// Exécuter la commande root
func Execute() error {
	return rootCmd.Execute()
}
