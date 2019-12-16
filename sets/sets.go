package sets

import (
	"fmt"

	"github.com/MagicTheGathering/mtg-sdk-go"
)

// FormattedMessage gets the full formatted version message
func ListarSets() {
	// qry := mtg.NewSetQuery()
	// qry.Where(mtg.CardSet, "ELD").OrderBy(mtg.CardNumber)
	sets, err := mtg.NewSetQuery().All()
	if err != nil {
		panic(err)
	}

	for _, set := range sets {
		fmt.Println(set)
	}

}
