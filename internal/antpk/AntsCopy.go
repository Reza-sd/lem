package antpk

func AntsCopy(baseAnts Ants, secondAnts *Ants) {

	//myAnts.AntsMap = make(map[int]Ant)
	secondAnts.AntsMap = make(map[int]Ant)
	secondAnts.AntsNumber = baseAnts.AntsNumber
	var ant Ant

	for key, value := range baseAnts.AntsMap {

		ant = secondAnts.AntsMap[key]
		ant.Name = value.Name
		ant.CurrentRoomName = value.CurrentRoomName
		ant.VisitedRoomsArr = []string{value.VisitedRoomsArr[0]}
		//make([]string, len(value.VisitedRoomsArr))

		secondAnts.AntsMap[key] = ant

	}

}
