/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

// lstDptCmd représente la commande lstDpt
var lstDptCmd = &cobra.Command{
	Use:   "lstDpt",
	Short: "Cette commande permet de lister tous les départements existants dans le fichier Source",
	Run: func(cmd *cobra.Command, args []string) {
		// Chargement des variables d'environnement
		err := godotenv.Load("C:\\Users\\Utilisateur\\Desktop\\MELIODAS\\meliodas_cli\\.env")
		if err != nil {
			fmt.Println("Erreur lors du chargement du fichier .env :", err)
			os.Exit(1)
		}

		// Récupération du chemin du fichier source
		sourceFile := os.Getenv("SOURCE_FILE")
		if sourceFile == "" {
			fmt.Println("La variable d'environnement SOURCE_FILE n'est pas définie.")
			os.Exit(1)
		}

		// Lecture du fichier CSV
		file, err := os.Open(sourceFile)
		if err != nil {
			fmt.Printf("Erreur lors de l'ouverture du fichier %s : %v\n", sourceFile, err)
			os.Exit(1)
		}
		defer file.Close()

		// Initialisation du lecteur
		r := csv.NewReader(file)
		r.Comma = ','

		// Utilisation d'un map pour éviter les doublons
		listeDepartements := make(map[string]struct{})

		// Lecture de chaque ligne du fichier
		for {
			record, err := r.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Printf("Erreur lors de la lecture de la ligne : %v\n", err)
				continue
			}

			// Initialisation de la colonne département
			if len(record) > 9 { // Assurez-vous qu'il y a suffisamment de colonnes
				departement := record[9][:2]
				listeDepartements[departement] = struct{}{} // Ajout dans le map
			}
		}

		// Affichage des départements
		fmt.Println("Liste des départements :")
		for dpt := range listeDepartements {
			fmt.Println(dpt)
		}
	},
}

func init() {
	RootCmd.AddCommand(lstDptCmd)
}
