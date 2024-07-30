package antpk

type Ant struct {
	Name            string
	CurrentRoomName string
	VisitedRoomsArr []string
	AssignedPath    []string
}

type Ants struct {
	AntsNumber int
	AntsMap    map[int]Ant
}

type TravelPlan struct {
	Steps    int
	SuccessfulPlan Ants
}