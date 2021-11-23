package entities

type Room struct {
	Type string
}

type Dungeon struct {
	Rooms []Room
}

func NewDungeon(howManyRooms int) *Dungeon {
	return &Dungeon{Rooms: make([]Room, howManyRooms)}
}
