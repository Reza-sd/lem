package move

type Ant struct {
	Name            string
	currentRoomName string
	VisitedRooms   []string
}

type Ants struct {
	Ants map[int]Ant
}
