package components

import (
	"fmt"
	"os"
	"time"
)

func LogWriter() {
	//Initialisation d'une variable now
	now := time.Now()

	// Ajout des informations aux fichier log
	logFilePath := os.Getenv("LOG_FILE")
	logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Une erreur s'est produite lors de l'ouverture du fichier log : %v", err)
		return
	}
	defer logFile.Close()
	logFile.WriteString(fmt.Sprintf("\nUne extraction a été effectuée le : %s | depuis le fichier source %s | vers nouveau fichier %v | les filtres actifs sont INSEE : %v et DPT :%v \n", now.Format("2006-01-02 15:04:05"), FilePath, ExtractFilePath+CompleteExtractFileName+".csv", CityINSEE, DepartID))
}
