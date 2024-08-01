package antpk

func AntsCopy(baseAnts AntGroup, secondAnts *AntGroup) {

	//myAnts.AntsMap = make(map[int]Ant)
	secondAnts.AntsMap = make(map[int]Ant)
	secondAnts.NumberOfAnts = baseAnts.NumberOfAnts
	//secondAnts.NumberOfSequence = baseAnts.NumberOfSequence

	var ant Ant

	for key, value := range baseAnts.AntsMap {

		ant = secondAnts.AntsMap[key]
		ant.Name = value.Name
		ant.StepNumber = value.StepNumber
		ant.CurrentRoomName = value.CurrentRoomName
		ant.VisitedRoomsArr = []string{value.VisitedRoomsArr[0]}
		//make([]string, len(value.VisitedRoomsArr))
		//copy(ant.VisitedRoomsArr,value.VisitedRoomsArr)
		secondAnts.AntsMap[key] = ant

	}

}
