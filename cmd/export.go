package cmd

import (
	"fmt"

	"github.com/abdealt/meliodas/components"
	"github.com/spf13/cobra"
)

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Cette commande sert à exporter les données.",
	Long:  `Cette commande sert à exporter les données. Grâce au fichier .env de Meliodas, on définit sur quoi on exporte, le fichier source et le lieu où seront extraites les données.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("La commande 'export' est éxécuté\n")
		fmt.Printf("Début du traitement \n")
		components.ReadCSVFileContentAndExtracter()
		fmt.Printf("Fin du traitement \n")
	},
}

func init() {
	rootCmd.AddCommand(exportCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// exportCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// exportCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
