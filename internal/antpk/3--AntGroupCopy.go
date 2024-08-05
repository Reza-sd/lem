package antpk

//import "fmt"

func AntGroupCopyAtFirstRoom(original AntGroup) (AntGroup, error) {
	funcName := "AntsCopy"
	//------------input validation-------------------
	var copy AntGroup
	if original.NumberOfAnts == 0 {
		return copy, logger.RWarnStr(funcName, "NumberOfAnts ? 0", "is not valid", "Empty AntsCopy")
	}

	//-----------------------------
	copy = AntGroup{
		NumberOfAnts:          original.NumberOfAnts,
		currentSequenceNumber: original.currentSequenceNumber,
		AntsDb:                make(map[string]Ant, len(original.AntsDb)),
		UsedTunnel:            map[int]map[string]string{},
	}

	for key, ant := range original.AntsDb {

		copy.AntsDb[key] = Ant{
			Name:            ant.Name,
			CurrentRoomName: ant.CurrentRoomName,
			VisitedRoomsArr: append([]string(nil), ant.VisitedRoomsArr...),
			StepNumber:      ant.StepNumber,
		}

	}
	return copy, nil
}

//==========================================
