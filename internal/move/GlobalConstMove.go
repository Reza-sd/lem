package move

type Ant struct {
	Name            string
	CurrentRoomName string
	VisitedRooms    []string
}

type Ants struct {
	//AntsNumber int
	Ants map[int]Ant
}
