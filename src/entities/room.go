package entities

const (
	NormalDungeon = 10
	//EpicDungeon = 20
	//InfiniteMaze = 100
)

type Room struct {
	Type string
}

type Dungeon struct {
	Rooms []Room
}

func NewDungeon(howManyRooms int) *Dungeon {
	return &Dungeon{Rooms: make([]Room, howManyRooms)}
}
