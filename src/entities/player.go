package entities

import (
	"fmt"
	"os"
)

const (
	MaxCardsOnHand = 5
)

type Player struct {
	Hand Hand
}

type Hand struct {
	Cards []Card
}

func newPlayer() *Player {
	return &Player{
		Hand: NewHand(),
	}
}

func NewHand() Hand {
	return Hand{
		Cards: []Card{},
	}
}

func NewPlayers(howManyPlayers int) *[]Player {
	var players []Player

	for index := 0; index < howManyPlayers; index++ {
		players = append(players, *newPlayer())
	}

	return &players
}

func (hand *Hand) KeepCard(drawnCard Card, discard *[]Card) {
	var choice int

	hand.Cards = append(hand.Cards, drawnCard)

	if len(hand.Cards) > MaxCardsOnHand {
		fmt.Printf("Sua mão contém %d cartas, você precisa descartar %d carta.\n", len(hand.Cards), len(hand.Cards)-MaxCardsOnHand)
		hand.listCards()
		fmt.Println("Escolha uma carta para descartar:")
		scanPlayerChoice(&choice)

		if choice <= len(hand.Cards) {
			discarded := hand.discard(choice)
			*discard = append(*discard, discarded)
		} else {
			exit()
		}
	}
}

func scanPlayerChoice(choice *int) {
	if _, err := fmt.Scanln(&choice); err != nil {
		fmt.Println(err)
		exit()
	}
}

func (hand *Hand) listCards() {
	var cardList string

	for index, card := range hand.Cards {
		if len(cardList) == 0 {
			cardList = fmt.Sprintf("%d-%v", index+1, card)
			continue
		}
		cardList += fmt.Sprintf(" | %d-%v", index+1, card)
	}

	fmt.Println(cardList)
}

func (hand *Hand) discard(position int) Card {
	realPosition := position - 1
	discardedItem := hand.Cards[realPosition]

	copy(hand.Cards[realPosition:], hand.Cards[realPosition+1:])
	hand.Cards[len(hand.Cards)-1] = Card{}
	hand.Cards = hand.Cards[:len(hand.Cards)-1]

	return discardedItem
}

func exit() {
	fmt.Println("Opção inválida, finalizando o programa")
	os.Exit(1)
}
