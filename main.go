package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/abdealt/meliodas_cli/cmd"
)

func main() {
	// Afficher le logo défini dans le fichier root.go lors du démarrage
	fmt.Println("Bienvenue dans Meliodas CLI!")

	// Créer un reader pour capturer l'entrée utilisateur en continu
	reader := bufio.NewReader(os.Stdin)

	// Boucle d'exécution continue
	for {
		// Demander une commande à l'utilisateur
		fmt.Print("\nEntrez une commande ('export', 'stat', 'lstdpt', 'help' ou 'exit' pour quitter) : ")
		input, _ := reader.ReadString('\n') // Lire l'entrée utilisateur
		input = strings.TrimSpace(input)    // Supprimer les espaces inutiles

		// Si l'utilisateur tape "exit", sortir du programme
		if input == "exit" {
			fmt.Println("Fermeture de l'application.")
			os.Exit(0)
		}

		// Diviser l'entrée utilisateur en arguments
		args := strings.Split(input, " ")

		// Passer les arguments à la commande root définie dans le fichier root.go
		cmd.RootCmd.SetArgs(args)

		// Exécuter la commande et gérer les erreurs
		if err := cmd.Execute(); err != nil {
			fmt.Printf("Erreur lors de l'exécution de la commande : %v\n", err)
		}
	}
}
