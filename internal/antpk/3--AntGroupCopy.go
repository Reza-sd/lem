package antpk

//import "fmt"

func AntGroupCopy(baseAntGroup AntGroup, secondAntGroup *AntGroup) error {
	funcName := "AntsCopy"
	//myAnts.AntsMap = make(map[int]Ant)
	//------------input validation-------------------

	if baseAntGroup.NumberOfAnts == 0 {
		return logger.RWarnStr(funcName, "NumberOfAnts ? 0", "is not valid", "Empty AntsCopy")
	}

	//-----------------------------
	secondAntGroup.AntsMap = make(map[int]Ant)
	secondAntGroup.NumberOfAnts = baseAntGroup.NumberOfAnts
	secondAntGroup.NumberOfSequence = baseAntGroup.NumberOfSequence

	var tempAnt Ant

	for key, ant := range baseAntGroup.AntsMap {

		tempAnt = secondAntGroup.AntsMap[key]
		tempAnt.Name = ant.Name
		tempAnt.StepNumber = ant.StepNumber
		tempAnt.CurrentRoomName = ant.CurrentRoomName

		tempAnt.VisitedRoomsArr = deepCopySlice(ant.VisitedRoomsArr)
		secondAntGroup.AntsMap[key] = tempAnt

	}
	return nil
}

// -------------
func deepCopySlice(original []string) []string {
	copy := make([]string, len(original))
	for i, item := range original {
		copy[i] = item
	}
	return copy
}

//-------------
