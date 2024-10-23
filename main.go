package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/abdealt/meliodas_cli/cmd"
)

func main() {
	fmt.Println("Bienvenue dans l'application CLI MELIODAS.")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("\nEntrez une commande ('export', 'stat', 'help', 'exit') : ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Vérification de l'entrée utilisateur
		switch input {
		case "exit":
			fmt.Println("Sortie de l'application.")
			return
		case "help":
			// Affichage de la desxription etc..
			cmd.Execute()
		default:
			// Simulation de l'appel de la commande avec Cobra
			args := strings.Split(input, " ")
			os.Args = append([]string{"meliodas"}, args...)
			cmd.Execute()
		}
	}
}
