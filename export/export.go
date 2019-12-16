package export

import (
	"fmt"

	"github.com/MagicTheGathering/mtg-sdk-go"
)

// FormattedMessage gets the full formatted version message
func ApresentaListagem() {
	qry := mtg.NewQuery()
	qry.Where(mtg.CardSet, "ELD").OrderBy(mtg.CardNumber)
	cards, err := qry.All()
	if err != nil {
		panic(err)
	}

	card := cards[0]

	fmt.Println(card.Name)

	// for _, card := range cards {
	// 	fmt.Println(card)
	// }
}
