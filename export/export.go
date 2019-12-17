package export

import (
	"fmt"

	"github.com/MagicTheGathering/mtg-sdk-go"
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

	for i, card := range cards {
		lines[i] = fmt.Sprintf("%s; %s; %s; %s; %s;\n", card.Number, card.Name, card.Type, card.ManaCost, card.Rarity)
	}

	file.Salvar(lines)
}
