package antpk

func AntsCopy(baseAntGroup AntGroup, secondAntGroup *AntGroup) error {
	funcName := "AntsCopy"
	//myAnts.AntsMap = make(map[int]Ant)
	//------------input validation-------------------

	if baseAntGroup.NumberOfAnts == 0 {
		return logger.RWarnStr(funcName, "NumberOfAnts ? 0", "is not valid", "Empty AntsCopy")
	}

	//-----------------------------
	secondAntGroup.AntsMap = make(map[int]Ant)
	secondAntGroup.NumberOfAnts = baseAntGroup.NumberOfAnts
	//secondAnts.NumberOfSequence = baseAnts.NumberOfSequence

	var ant Ant

	for key, value := range baseAntGroup.AntsMap {

		ant = secondAntGroup.AntsMap[key]
		ant.Name = value.Name
		ant.StepNumber = value.StepNumber
		ant.CurrentRoomName = value.CurrentRoomName
		ant.VisitedRoomsArr = []string{value.VisitedRoomsArr[0]}
		//make([]string, len(value.VisitedRoomsArr))
		//copy(ant.VisitedRoomsArr,value.VisitedRoomsArr)
		secondAntGroup.AntsMap[key] = ant

	}
	return nil
}
