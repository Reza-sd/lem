package antpk

//import "fmt"

func AntGroupCopyAtFirstRoom(baseAntGroup AntGroup, secondAntGroup *AntGroup) error {
	funcName := "AntsCopy"
	//myAnts.AntsMap = make(map[int]Ant)
	//------------input validation-------------------

	if baseAntGroup.NumberOfAnts == 0 {
		return logger.RWarnStr(funcName, "NumberOfAnts ? 0", "is not valid", "Empty AntsCopy")
	}

	//-----------------------------
	secondAntGroup.AntsMap = map[int]Ant{}
	secondAntGroup.NumberOfAnts = baseAntGroup.NumberOfAnts
	secondAntGroup.SequenceNumber = 0
	secondAntGroup.UsedTunnel.UsedTunnelsMap = map[int]map[string]string{}
	//-----------------------------
	//var tempAnt Ant

	for key, ant := range baseAntGroup.AntsMap {

		secondAntGroup.AntsMap[key]=Ant{
			Name: ant.Name,
			CurrentRoomName: ant.CurrentRoomName,
			VisitedRoomsArr: append([]string(nil), ant.VisitedRoomsArr...),
			StepNumber:      ant.StepNumber,

		}

		// tempAnt = secondAntGroup.AntsMap[key]
		// tempAnt.Name = ant.Name
		// tempAnt.StepNumber = ant.StepNumber
		// tempAnt.CurrentRoomName = ant.CurrentRoomName
		// tempAnt.VisitedRoomsArr = deepCopySlice(ant.VisitedRoomsArr)
		// secondAntGroup.AntsMap[key] = tempAnt

	}
	return nil
}

// -------------
// func deepCopySlice(original []string) []string {
// 	copy := make([]string, len(original))
// 	for i, item := range original {
// 		copy[i] = item
// 	}
// 	return copy
// }

//-------------
