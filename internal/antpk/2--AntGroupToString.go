package antpk

import (
	"fmt"
)

func (allAnts *AntGroup) ToString() (string, error) {

	var antGroupString string
	antGroupString = fmt.Sprintf("NumberOfAnts=%v, SequenceNumber=%v, UsedTunnel=%v \n", allAnts.NumberOfAnts, allAnts.CurrentSequenceNumber, allAnts.UsedTunnel)
	count := allAnts.NumberOfAnts
	if count == 0 {
		antGroupString = "<< NO ANT,Empty AntGroup >>"
		return antGroupString, nil
	}

	for i := 1; i <= count; i++ {

		ant := allAnts.AntsDb[fmt.Sprintf("L%v", i)]
		antStr := fmt.Sprintf("antName=%v, CurrentRoom=%v, StepNumber=%v, VisitedRooms=%v", ant.Name, ant.CurrentRoomName, ant.StepNumber, ant.VisitedRoomsArr)
		antGroupString += antStr
		if i != allAnts.NumberOfAnts {
			antGroupString += "\n"
		}

	}

	return antGroupString, nil

}

// --------------------
func (allAnts *AntGroup) Print() {
	fmt.Println("\n----------AntGroup------------")
	str, _ := allAnts.ToString()
	fmt.Println(str)
	fmt.Println("---------------------------")
}

//-------------------------
