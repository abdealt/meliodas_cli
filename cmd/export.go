package cmd

import (
	"fmt"

	"github.com/abdealt/meliodas/components"
	"github.com/spf13/cobra"
)

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Cette commande sert a exporter les données.",
	Long:  `Cette commande sert a exporter les données. Grace au fichie .env de Meliodas, on definit sur quoi on export, le fichier source et le lieu ou seront extraite les données.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Début du traitement")
		components.ReadCSVFileContentAndExtracter()
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
