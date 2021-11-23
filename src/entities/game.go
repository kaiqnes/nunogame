package entities

type Game struct {
	Players            []Player
	Dungeon            Dungeon
	Doors              *DoorDeck
	Treasures          TreasureDeck
	DiscardedDoors     []Card
	DiscardedTreasures []Card
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
