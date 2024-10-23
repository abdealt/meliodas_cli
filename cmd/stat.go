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
		fmt.Println("La commande stat est éxécutée.")
		// Logique pour afficher les statistiques
		fmt.Printf("Il y a %v éléments totaux dans le fichier source. Sur tous ces éléments, il y a %v éléments exportés.", components.ComptTotal, components.ComptElement)
	},
}

// Fonction lors de l'appel de la commande
func ExecuteStat() error {
	return statCmd.Execute()
}

func init() {
	RootCmd.AddCommand(statCmd)
}
