package antpk

import (
	"fmt"
)

func (allAnts *AntGroup) ToString() (string, error) {

	var antGroupString string
	count := allAnts.NumberOfAnts
	if count == 0 {
		antGroupString = "<< NO ANT,Empty AntGroup >>"
		return antGroupString, nil
	}

	for i := 1; i <= count; i++ {

		ant := allAnts.AntsMap[i]
		antStr := fmt.Sprintf("ant=%v, CurrentRoom=%v, StepNumber=%v, VisitedRooms=%v", ant.Name, ant.CurrentRoomName, ant.StepNumber, ant.VisitedRoomsArr)
		antGroupString += antStr
		if i != allAnts.NumberOfAnts {
			antGroupString += "\n"
		}
	}

	return antGroupString, nil

}
