package cmd

import (
	"fmt"

	"github.com/abdealt/meliodas/components"
	"github.com/spf13/cobra"
)

// statCmd est la commande stat ici les infos et propriétés
var statCmd = &cobra.Command{
	Use:   "stat",
	Short: "Cette commande sert à afficher les statistiques obtenues lors du dernier traitement.",
	Run: func(cmd *cobra.Command, args []string) {
		// Affichage de l'état
		fmt.Printf("Commande 'stat' exécutée\n")

		// Affichage des statistiques
		components.GetStats() // Récupérer les statistiques
	},
}

// Fonction lors de l'appel de la commande
func ExecuteStat() error {
	return statCmd.Execute()
}

func init() {
	rootCmd.AddCommand(statCmd)
}
