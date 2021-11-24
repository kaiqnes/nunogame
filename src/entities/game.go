package entities

import "fmt"

type Game struct {
	Players            []Player
	Dungeon            Dungeon
	Doors              *DoorDeck
	Treasures          TreasureDeck
	DiscardedDoors     []Card
	DiscardedTreasures []Card
}

func (game *Game) KickTheDoor(player *Player) {
	// How many cards should draw?
	cards := game.Doors.DrawCards(quantityToDraw(player.Level))
	fmt.Printf("%s chutou a porta e encontrou %v\n", player.Name, cards)

	handleWithCardsFromKickedDoor(player, cards)
}

func handleWithCardsFromKickedDoor(player *Player, cards []Card) {
	var (
		remainingCards []Card
		cardCommands   = map[string]interface{}{CardTypeCurse: "ApplyCurseCommand", CardTypeClass: "KeepClass", CardTypeMonster: "StartBattle", CardTypeRace: "KeepRace"}
	)

	for _, card := range cards {
		if card.Type == CardTypeClass || card.Type == CardTypeRace {
			remainingCards = append(remainingCards, card)
			continue
		}

		commandToExecute := cardCommands[card.Type]
		fmt.Printf("Executing commang %s\n", commandToExecute)
	}

	if len(remainingCards) > 0 {
		for _, card := range remainingCards {
			player.Hand.Cards = append(player.Hand.Cards, card)
		}
	}
}

func NewGame(howManyPlayers, howManyRooms int) *Game {
	game := Game{
		Players:            *NewPlayers(howManyPlayers),
		Dungeon:            *NewDungeon(howManyRooms),
		Doors:              NewDoorDeck(),
		Treasures:          TreasureDeck{},
		DiscardedDoors:     []Card{},
		DiscardedTreasures: []Card{},
	}

	game.Doors.ShuffleDeck()
	return &game
}

func quantityToDraw(playerLevel uint8) (howManyCardsShouldDraw uint8) {
	howManyCardsShouldDraw = playerLevel / 10
	if howManyCardsShouldDraw == 0 {
		howManyCardsShouldDraw = 1
	}
	return
}
