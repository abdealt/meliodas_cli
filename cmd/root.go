package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/abdealt/meliodas/csvExtracter"
	"github.com/abdealt/meliodas_cli/assets"
	"github.com/common-nighthawk/go-figure"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

// Variables globales
var (
	TotalElements     int
	ExtractedElements int
	my_WI             *csvExtracter.WorkerImmeuble
	err               error
	envPath           string // Nouveau flag pour le chemin du fichier .env
)

// Définir la commande root
var RootCmd = &cobra.Command{
	Use:   "meliodas",
	Short: "\nUn CLI pour lire un fichier",
	Long:  "\nMeliodas CLI est un outil en ligne de commande simple et efficace conçu pour automatiser le traitement de fichiers CSV...",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Charger le fichier .env à partir du chemin spécifié ou d'un chemin par défaut
		if envPath == "" {
			envPath = "C:\\Users\\Utilisateur\\Desktop\\MELIODAS\\meliodas_cli\\.env"
		}
		if err := godotenv.Load(envPath); err != nil {
			fmt.Println("Erreur lors du chargement du fichier .env:", err)
			os.Exit(1)
		}

		// Initialiser la configuration depuis le fichier .env
		var mycfg csvExtracter.Config
		mycfg.File_immeuble = os.Getenv("SOURCE_FILE")
		mycfg.File_export = os.Getenv("EXTRACT_FILE")
		mycfg.File_log = os.Getenv("LOG_FILE")

		// Traitement des départements
		listeDept := strings.TrimSpace(os.Getenv("DEPARTMENT_ID"))
		if listeDept != "" {
			mycfg.Lst_Dprt = strings.Split(listeDept, ",")
			for i := range mycfg.Lst_Dprt {
				mycfg.Lst_Dprt[i] = strings.TrimSpace(mycfg.Lst_Dprt[i])
			}
		}

		// Créer le WorkerImmeuble
		my_WI, err = csvExtracter.NewWorkerImmeuble(mycfg)
		if err != nil {
			fmt.Println("Erreur lors de la création de WorkerImmeuble:", err.Error())
			return
		}
	},
}

func init() {
	// Ajout du flag pour le fichier .env
	RootCmd.PersistentFlags().StringVar(&envPath, "env", "", "Chemin vers le fichier .env")

	myLogo := figure.NewFigure(assets.AppDisplayName, "", true)
	myLogo.Print()
}

// Exécuter la commande root
func Execute() error {
	return RootCmd.Execute()
}
