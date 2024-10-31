package main

import (
	"fmt"
	"os"

	"github.com/abdealt/meliodas_cli/cmd"
)

func main() {
	// Afficher le logo défini dans le fichier root.go lors du démarrage
	fmt.Println("Bienvenue dans Meliodas CLI!")

	// Passer les arguments de la ligne de commande à la commande root
	cmd.RootCmd.SetArgs(os.Args[1:])

	// Exécuter la commande et gérer les erreurs
	if err := cmd.Execute(); err != nil {
		fmt.Printf("Erreur lors de l'exécution de la commande : %v\n", err)
		os.Exit(1) // Sortir avec un code d'erreur si l'exécution échoue
	}
}
