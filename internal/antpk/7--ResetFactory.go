package antpk

import (
// "fmt"
// "fmt"
// "fmt"
// "main/internal/graphpk"
)

func (myAntGroup *AntGroup) ResetFactory(currnetRoomName string) {
	//currnetRoomName:=
	myAntGroup.CurrentSequenceNumber = 0
	myAntGroup.NotArrivedAntsName = make(map[string]struct{})
	myAntGroup.UsedTunnel = make(map[string][]string)

	for antName, _ := range myAntGroup.AntsDb {

		//antObject.CurrentRoomName=currnetRoomName
		myAntGroup.NotArrivedAntsName[antName] = struct{}{}

		ant := myAntGroup.AntsDb[antName]
		ant.CurrentRoomName = currnetRoomName
		ant.StepNumber = 0
		ant.VisitedRoomsArr = []string{currnetRoomName}
		myAntGroup.AntsDb[antName] = ant
	}

}
