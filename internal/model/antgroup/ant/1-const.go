package ant

type Mta = uint16 //my type ant pk

type Ant struct {
	Name            Mta
	CurrentRoomName Mta
	StepNumber      Mta
	VisitedRoomsArr []Mta
}
