package cmd

import (
	"fmt"

	"github.com/abdealt/meliodas/components"
	"github.com/spf13/cobra"
)

// exportCmd Est la commande exort ici les infos et propriétés
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Cette commande sert à exporter les données.",
	Long:  `Cette commande sert à exporter les données. Grâce au fichier .env de Meliodas, on définit sur quoi on exporte, le fichier source et le lieu où seront extraites les données.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Traitement en cours...")
		components.ReadCSVFileContentAndExtracter()

		// Simuler l'ajout d'éléments
		TotalElements += 100    // Exemple
		ExtractedElements += 50 // Exemple

		fmt.Println("Fin du traitement")
	},
}

func ExecuteExport() error {
	return exportCmd.Execute()
}

func init() {
	rootCmd.AddCommand(exportCmd)
}
