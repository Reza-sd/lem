package move

type Ant struct {
	Name            string
	CurrentRoomName string
	VisitedRoomsArr []string
	AssignedPath    []string
}

type Ants struct {
	AntsMap    map[int]Ant
	AntsNumber int
}
