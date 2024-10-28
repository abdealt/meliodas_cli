package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/abdealt/meliodas/components"
	"github.com/abdealt/meliodas_cli/assets"
	"github.com/common-nighthawk/go-figure"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

// Variables globales
var TotalElements int
var ExtractedElements int
var my_WI *components.WorkerImmeuble
var err error

// Définir la commande root
var RootCmd = &cobra.Command{
	Use:   "meliodas",
	Short: "\nUn CLI pour lire un fichier",
	Long:  "\nMeliodas CLI est un outil en ligne de commande simple et efficace conçu pour automatiser le traitement de fichiers CSV.\nIl permet aux utilisateurs d'extraire rapidement des données spécifiques à partir de fichiers, sans avoir à ouvrir un logiciel complexe.\n\nSi vous souahitez changer d'opération, il faut changer le fichier de configuration, pour spécifier le (ou les) DÉPARTEMENT(S) de votre choix, ou le (ou les) codes INSEE(S) de votre choix.\n\n",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Chargement du fichier .ENV
		er := godotenv.Load("C:\\Users\\Utilisateur\\Desktop\\MELIODAS\\meliodas_cli\\.env")
		if er != nil {
			fmt.Println("Erreur lors du chargement du fichier.env")
			os.Exit(1)
		}
		// Initialiser la configuration depuis le fichier .ENV
		var mycfg components.Config
		mycfg.File_immeuble = os.Getenv("SOURCE_FILE")
		mycfg.File_export = os.Getenv("EXTRACT_FILE")
		mycfg.File_log = os.Getenv("LOG_FILE")

		listeInseeAvant := os.Getenv("CITY_INSEE")
		listeInseeApres := strings.Split(listeInseeAvant, ",")
		listeDeptAvant := os.Getenv("DEPARTMENT_ID")
		listeDeptApres := strings.Split(listeDeptAvant, ",")
		mycfg.Lst_Insee = listeInseeApres
		mycfg.Lst_Dprt = listeDeptApres

		// Vérification des Informations
		if mycfg.File_immeuble == "" {
			fmt.Println("Veuillez spécifier un fichier source dans la configuration")
		}
		if mycfg.File_export == "" {
			fmt.Println("Veuillez spécifier un chemin pour exporter dans la configuration")
		}
		if len(mycfg.Lst_Dprt) == 0 && len(mycfg.Lst_Insee) == 0 {
			fmt.Println("Veuillez spécifier au moins un code INSEE ou un département dans la configuration")
		}

		// Créer le WorkerImmeuble
		my_WI, err = components.NewWorkerImmeuble(mycfg)
		if err != nil {
			fmt.Println("Erreur lors de la création de WorkerImmeuble:", err.Error())
			return
		}
	},
}

func init() {
	myLogo := figure.NewFigure(assets.AppDisplayName, "", true)
	myLogo.Print()
}

// Exécuter la commande root
func Execute() error {
	return RootCmd.Execute()
}
