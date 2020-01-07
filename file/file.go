package file

import (
	"io"
	"log"
	"os"
)

func Save(lista []string) {
	csvfile, err := os.Create("planilha.csv")

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	for _, card := range lista {
		_, err = io.WriteString(csvfile, card)
		if err != nil {
			log.Fatal(err)
		}
		csvfile.Sync()
	}

	defer csvfile.Close()

}
