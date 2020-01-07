// Package Card implements common functionality needed for sheet.
package card

import "fmt"

// This is the type that represents a Social Media Post
type Card struct {
	Number   string
	Name     string
	TypeCard string
	ManaCost string
	Rarity   string
}

// This is the function responsable for creating a new social media post
func NewCard(number, name, typeCard, manaCost, rarity string) *Card {
	return &Card{Number: number, Name: name, TypeCard: typeCard, ManaCost: manaCost, Rarity: rarity}
}

func (c Card) String() string {
	return fmt.Sprintf("%v; %v; %v; %v; %v;\n", c.Number, c.Name, c.TypeCard, c.ManaCost, c.Rarity)
}
