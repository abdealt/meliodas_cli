package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// statCmd est la commande stat ici les infos et propriétés
var statCmd = &cobra.Command{
	Use:   "stat",
	Short: "Cette commande sert à afficher les statistiques obtenues lors du dernier traitement.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Traitement en cours...")
		// if err := my_WI.ExtractStatisticsFromCSV(); err != nil {
		// 	fmt.Println("Erreur lors du traitement:", err)
		// 	return
		// }
		fmt.Println("Fin du traitement")
	},
}

// Fonction lors de l'appel de la commande
func ExecuteStat() error {
	return statCmd.Execute()
}

func init() {
	RootCmd.AddCommand(statCmd)
}
