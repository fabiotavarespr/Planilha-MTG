package export

import (
	"fmt"
	"github.com/MagicTheGathering/mtg-sdk-go"
	"github.com/fabiotavarespr/Planilha-MTG/card"
	"github.com/fabiotavarespr/Planilha-MTG/file"
)

// FormattedMessage gets the full formatted version message
func ApresentaListagem(set string) {
	fmt.Println("Starting find collection " + set)
	qry := mtg.NewQuery()
	qry.Where(mtg.CardSet, set).OrderBy(mtg.CardNumber)
	cards, err := qry.All()
	fmt.Println("Finishing find collection " + set)
	if err != nil {
		panic(err)
	}

	lines := make([]string, 0)

	sizeList := len(cards)

	for i := 0; i < sizeList; i++ {
		c1 := cards[i]
		// Logica criada para atender edições com flips ou adventures
		if i != sizeList-1 {
			c2 := cards[i+1]
			if c1.Number == c2.Number {
				c1.Name += " / " + c2.Name
				c1.Type += " / " + c2.Type
				i++
			}
		}
		lines = append(lines, card.NewCard(c1.Number, c1.Name, c1.Type, c1.ManaCost, c1.Rarity).String())
	}

	file.Save(set, lines)
}
