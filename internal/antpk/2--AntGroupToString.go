package antpk

import (
	"fmt"
)

func (allAnts *AntGroup) ToString() (string, error) {

	var antGroupString string
	antGroupString = fmt.Sprintf("NumberOfAnts=%v\nSequenceNumber=%v\nUsedTunnel=%v \nNotArrivedAntsName=%v\n", allAnts.NumberOfAnts, allAnts.CurrentSequenceNumber, allAnts.UsedTunnel, allAnts.NotArrivedAntsName)

	count := allAnts.NumberOfAnts
	if count == 0 {
		antGroupString = "<< NO ANT,Empty AntGroup >>"
		return antGroupString, nil
	}

	for i := mta(1); i <= count; i++ {

		//ant := allAnts.AntsDb[fmt.Sprintf("L%v", i)]
		ant := allAnts.AntsDb[i]
		antStr := fmt.Sprintf("%v-antName=%v, CurrentRoom=%v, StepNumber=%v, VisitedRooms=%v", i, ant.Name, ant.CurrentRoomName, ant.StepNumber, ant.VisitedRoomsArr)
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
