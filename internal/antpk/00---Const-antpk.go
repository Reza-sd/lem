package antpk

type Ant struct {
	Name            string
	CurrentRoomName string
	VisitedRoomsArr []string
	StepNumber      int
	//AssignedPath    []string
}

type Ants struct {
	NumberOfAnts     int
	AntsMap          map[int]Ant
	NumberOfSequence int
}

type TravelPlan struct {
	FinalSequence int
	TheBestPlan   Ants
}
