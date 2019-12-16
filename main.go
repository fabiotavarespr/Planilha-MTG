package main

import (
	"fmt"

	mtg "github.com/MagicTheGathering/mtg-sdk-go"
)

//import "github.com/fabiotavarespr/Planilha-MTG/cmd"

func main() {
	qry := mtg.NewQuery()
	fmt.Println(qry)
	qry.Where(mtg.CardSet, "ELD").OrderBy(mtg.CardNumber)
	fmt.Println(qry)
	cards, err := qry.All()
	fmt.Println(cards)
	fmt.Println(err)

	//	cmd.Execute()
}
