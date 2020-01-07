package export

import (
	"github.com/MagicTheGathering/mtg-sdk-go"
	"github.com/fabiotavarespr/Planilha-MTG/card"
	"github.com/fabiotavarespr/Planilha-MTG/file"
)

// FormattedMessage gets the full formatted version message
func ApresentaListagem() {
	qry := mtg.NewQuery()
	qry.Where(mtg.CardSet, "ELD").OrderBy(mtg.CardNumber)
	cards, err := qry.All()
	if err != nil {
		panic(err)
	}

	lines := make([]string, len(cards))

	for i, c := range cards {

		var name string

		if len(c.Names) > 1 {
			for i, n := range c.Names {
				if i == 0 {
					name = n
				} else {
					name = name + " / " + n
				}
			}
		} else {
			name = c.Name
		}
		typeCard := c.Type

		check := true

		if check {
			lines[i] = card.NewCard(c.Number, name, typeCard, c.ManaCost, c.Rarity).String()
		}

	}

	file.Save(lines)
}
