package entities

import (
	"math/rand"
	"time"
)

const (
	CardTypeCurse   = "CURSE"
	CardTypeMonster = "MONSTER"
	CardTypeClass   = "CLASS"
	CardTypeRace    = "RACE"
)

var (
	CARD1  = Card{ID: 1, Name: "Card 1", Type: CardTypeCurse}
	CARD2  = Card{ID: 2, Name: "Card 2", Type: CardTypeMonster}
	CARD3  = Card{ID: 3, Name: "Card 3", Type: CardTypeClass}
	CARD4  = Card{ID: 4, Name: "Card 4", Type: CardTypeRace}
	CARD5  = Card{ID: 5, Name: "Card 5", Type: CardTypeCurse}
	CARD6  = Card{ID: 6, Name: "Card 6", Type: CardTypeMonster}
	CARD7  = Card{ID: 7, Name: "Card 7", Type: CardTypeClass}
	CARD8  = Card{ID: 8, Name: "Card 8", Type: CardTypeRace}
	CARD9  = Card{ID: 9, Name: "Card 9", Type: CardTypeCurse}
	CARD10 = Card{ID: 10, Name: "Card 10", Type: CardTypeMonster}
	CARD11 = Card{ID: 11, Name: "Card 11", Type: CardTypeClass}
	CARD12 = Card{ID: 12, Name: "Card 12", Type: CardTypeRace}
	CARD13 = Card{ID: 13, Name: "Card 13", Type: CardTypeCurse}
	CARD14 = Card{ID: 14, Name: "Card 14", Type: CardTypeMonster}
	CARD15 = Card{ID: 15, Name: "Card 15", Type: CardTypeClass}
	CARD16 = Card{ID: 16, Name: "Card 16", Type: CardTypeRace}
	CARD17 = Card{ID: 17, Name: "Card 17", Type: CardTypeCurse}
	CARD18 = Card{ID: 18, Name: "Card 18", Type: CardTypeMonster}
	CARD19 = Card{ID: 19, Name: "Card 19", Type: CardTypeClass}
	CARD20 = Card{ID: 20, Name: "Card 20", Type: CardTypeRace}
	//CARD21 = Card{ID: 21, Name: "Card 21", Type: CardTypeCurse }
	//CARD22 = Card{ID: 22, Name: "Card 22", Type: CardTypeMonster }
	//CARD23 = Card{ID: 23, Name: "Card 23", Type: CardTypeClass }
	//CARD24 = Card{ID: 24, Name: "Card 24", Type: CardTypeRace }
	//CARD25 = Card{ID: 25, Name: "Card 25", Type: CardTypeCurse }
	//CARD26 = Card{ID: 26, Name: "Card 26", Type: CardTypeMonster }
	//CARD27 = Card{ID: 27, Name: "Card 27", Type: CardTypeClass }
	//CARD28 = Card{ID: 28, Name: "Card 28", Type: CardTypeRace }
	//CARD29 = Card{ID: 29, Name: "Card 29", Type: CardTypeCurse }
	//CARD30 = Card{ID: 30, Name: "Card 30", Type: CardTypeMonster }
	//CARD31 = Card{ID: 31, Name: "Card 31", Type: CardTypeClass }
	//CARD32 = Card{ID: 32, Name: "Card 32", Type: CardTypeRace }
	//CARD33 = Card{ID: 33, Name: "Card 33", Type: CardTypeCurse }
	//CARD34 = Card{ID: 34, Name: "Card 34", Type: CardTypeMonster }
	//CARD35 = Card{ID: 35, Name: "Card 35", Type: CardTypeClass }
	//CARD36 = Card{ID: 36, Name: "Card 36", Type: CardTypeRace }
	//CARD37 = Card{ID: 37, Name: "Card 37", Type: CardTypeCurse }
	//CARD38 = Card{ID: 38, Name: "Card 38", Type: CardTypeMonster }
	//CARD39 = Card{ID: 39, Name: "Card 39", Type: CardTypeClass }
	//CARD40 = Card{ID: 40, Name: "Card 40", Type: CardTypeRace }
	//CARD41 = Card{ID: 41, Name: "Card 41", Type: CardTypeCurse }
	//CARD42 = Card{ID: 42, Name: "Card 42", Type: CardTypeMonster }
	//CARD43 = Card{ID: 43, Name: "Card 43", Type: CardTypeClass }
	//CARD44 = Card{ID: 44, Name: "Card 44", Type: CardTypeRace }
	//CARD45 = Card{ID: 45, Name: "Card 45", Type: CardTypeCurse }
	//CARD46 = Card{ID: 46, Name: "Card 46", Type: CardTypeMonster }
	//CARD47 = Card{ID: 47, Name: "Card 47", Type: CardTypeClass }
	//CARD48 = Card{ID: 48, Name: "Card 48", Type: CardTypeRace }
	//CARD49 = Card{ID: 49, Name: "Card 49", Type: CardTypeCurse }
	//CARD50 = Card{ID: 50, Name: "Card 50", Type: CardTypeMonster }
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
	Type string
}

func NewDoorDeck() *DoorDeck {
	return &DoorDeck{
		Cards: []Card{
			CARD1,
			CARD2,
			CARD3,
			CARD4,
			CARD5,
			CARD6,
			CARD7,
			CARD8,
			CARD9,
			CARD10,
			CARD10,
			CARD11,
			CARD12,
			CARD13,
			CARD14,
			CARD15,
			CARD16,
			CARD17,
			CARD18,
			CARD19,
			CARD20,
		},
	}
}

func (deck *DoorDeck) DrawCard() (card Card) {
	return deck.pop()
}

func (deck DoorDeck) ShuffleDeck() *DoorDeck {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck.Cards), func(i, j int) {
		deck.Cards[i], deck.Cards[j] = deck.Cards[j], deck.Cards[i]
	})
	return &deck
}

func (deck *DoorDeck) pop() Card {
	poppedElement := deck.Cards[0]
	smallerDeck := append(deck.Cards[:0], deck.Cards[1:]...)
	deck.Cards = smallerDeck
	return poppedElement
}

func (deck *DoorDeck) DiscardCard(card Card, discardedCards *[]Card) {
	*discardedCards = append(*discardedCards, card)
}

func (deck *DoorDeck) DrawCards(howManyCards uint8) (drawnCards []Card) {
	for index := uint8(0); index < howManyCards; {
		drawnCards = append(drawnCards, deck.pop())
		index++
	}
	return
}
