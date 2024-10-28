package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// exportCmd Est la commande export ici ses informations et ses propriétés
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Cette commande sert à exporter les données.",
	Long:  `Cette commande sert à exporter les données. Grâce au fichier .env de Meliodas, on définit sur quoi on exporte, le fichier source et le lieu où seront extraites les données.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Traitement en cours...")
		if err := my_WI.SuperreaderCSV(); err != nil {
			fmt.Println("Erreur lors du traitement:", err)
			return
		}
		fmt.Println("Fin du traitement")
	},
}

func ExecuteExport() error {
	return exportCmd.Execute()
}

func init() {
	RootCmd.AddCommand(exportCmd)
}
