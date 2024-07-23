package move

type Ant struct {
	Name            string
	CurrentRoomName string
	VisitedRooms    []string
}

type Ants struct {
	Ants map[int]Ant
}
