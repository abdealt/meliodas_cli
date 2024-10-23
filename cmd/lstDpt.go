/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

func ListeDept() {
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
		if len(record) >= 18 {
			if record[8] == "" {
				// Ligne incomplète, on passe à la suivante
				continue
			} else {
				codeDept := strings.TrimSpace(record[8][:2])
				// Ajout dans le map
				listeDepartements[codeDept] = struct{}{}
			}

		}
	}

	// Affichage des départements
	fmt.Printf("Liste des départements :\n")
	for dpt := range listeDepartements {
		fmt.Print(dpt + "|")
	}
}

// lstDptCmd représente la commande lstDpt
var lstDptCmd = &cobra.Command{
	Use:   "lstdpt",
	Short: "Cette commande permet de lister tous les départements existants dans le fichier Source",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Recherche des département dans le fichier source...")
		// Appel de la fonction pour lister les départements
		ListeDept()
	},
}

// Fonction lors de l'appel de la commande
func ExecuteLstDpt() error {
	return lstDptCmd.Execute()
}

func init() {
	RootCmd.AddCommand(lstDptCmd)
}
