package entities

import (
	"math/rand"
	"time"
)

var (
	DOOR1  = Card{ID: 1, Name: "Card 1"}
	DOOR2  = Card{ID: 2, Name: "Card 2"}
	DOOR3  = Card{ID: 3, Name: "Card 3"}
	DOOR4  = Card{ID: 4, Name: "Card 4"}
	DOOR5  = Card{ID: 5, Name: "Card 5"}
	DOOR6  = Card{ID: 6, Name: "Card 6"}
	DOOR7  = Card{ID: 7, Name: "Card 7"}
	DOOR8  = Card{ID: 8, Name: "Card 8"}
	DOOR9  = Card{ID: 9, Name: "Card 9"}
	DOOR10 = Card{ID: 10, Name: "Card 10"}
	DOOR11 = Card{ID: 11, Name: "Card 11"}
	DOOR12 = Card{ID: 12, Name: "Card 12"}
	DOOR13 = Card{ID: 13, Name: "Card 13"}
	DOOR14 = Card{ID: 14, Name: "Card 14"}
	DOOR15 = Card{ID: 15, Name: "Card 15"}
	DOOR16 = Card{ID: 16, Name: "Card 16"}
	DOOR17 = Card{ID: 17, Name: "Card 17"}
	DOOR18 = Card{ID: 18, Name: "Card 18"}
	DOOR19 = Card{ID: 19, Name: "Card 19"}
	DOOR20 = Card{ID: 20, Name: "Card 20"}
)

type Deck interface {
	TakeCard()
	ShuffleDeck()
	pop()
	DiscardCard(card Card, discardedCards []Card)
}

type DoorDeck struct {
	Cards []Card
}

type TreasureDeck struct {
	Cards []Card
}

type Card struct {
	ID   uint8
	Name string
}

func NewDoorDeck() *DoorDeck {
	return &DoorDeck{
		Cards: []Card{
			DOOR1,
			DOOR2,
			DOOR3,
			DOOR4,
			DOOR5,
			DOOR6,
			DOOR7,
			DOOR8,
			DOOR9,
			DOOR10,
			DOOR10,
			DOOR11,
			DOOR12,
			DOOR13,
			DOOR14,
			DOOR15,
			DOOR16,
			DOOR17,
			DOOR18,
			DOOR19,
			DOOR20,
		},
	}
}

func (deck *DoorDeck) DrawCard() (card Card) {
	return deck.pop(0)
}

func (deck DoorDeck) ShuffleDeck() *DoorDeck {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck.Cards), func(i, j int) {
		deck.Cards[i], deck.Cards[j] = deck.Cards[j], deck.Cards[i]
	})
	return &deck
}

func (deck *DoorDeck) pop(index int) Card {
	poppedElement := deck.Cards[index]
	smallerDeck := append(deck.Cards[:index], deck.Cards[index+1:]...)
	deck.Cards = smallerDeck
	return poppedElement
}

func (deck *DoorDeck) DiscardCard(card Card, discardedCards *[]Card) {
	*discardedCards = append(*discardedCards, card)
}
