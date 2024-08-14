package ant

type Mta = uint16 //my type ant pk

type Ant struct {
	//Name            mta
	CurrentRoomName Mta
	VisitedRoomsArr []Mta
	StepNumber      Mta
}