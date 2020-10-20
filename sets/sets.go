package sets

import (
	"fmt"

	"github.com/fabiotavarespr/mtg-sdk-go"
)

// FormattedMessage gets the full formatted version message
func ListarSets() {
	// qry := mtg.NewSetQuery()
	// qry.Where(mtg.CardSet, "ELD").OrderBy(mtg.CardNumber)
	sets, err := mtg.NewSetQuery().All()
	if err != nil {
		panic(err)
	}

	for _, s := range sets {
		fmt.Printf("%s - %s\n", s.SetCode, s.Name)
	}

}
