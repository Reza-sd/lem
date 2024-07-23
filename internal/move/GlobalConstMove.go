package move

type Ant struct {
	Name            string
	currentRoomName string
}

type Ants struct {
	Ants map[int]Ant
}
