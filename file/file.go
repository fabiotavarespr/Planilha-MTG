package file

import (
	"encoding/csv"
	"log"
	"os"
)

func Salvar(row []string) {
	csvfile, err := os.Create("planilha.csv")

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	csvwriter := csv.NewWriter(csvfile)

	csvwriter.Write(row)

	csvwriter.Flush()

	csvfile.Close()
}
