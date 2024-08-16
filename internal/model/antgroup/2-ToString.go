package antgroup

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

	for i := Mtag(1); i <= count; i++ {

		ant := allAnts.AntsDb[i]

		antStr, _ := ant.ToString()
		antGroupString += fmt.Sprintf("\n%v-%v", i, antStr)

	}

	return antGroupString, nil

}
