package antpk

import (
// "fmt"
// "fmt"
// "fmt"
// "main/internal/graphpk"
)

func (myAntGroup *AntGroup) ResetFactory() {
	//currnetRoomName:=
	myAntGroup.CurrentSequenceNumber = 0
	myAntGroup.NotArrivedAntsName = make(map[mt]struct{})
	myAntGroup.UsedTunnel = make(map[mt][]mt)

	for antName := range myAntGroup.AntsDb {

		//antObject.CurrentRoomName=currnetRoomName
		myAntGroup.NotArrivedAntsName[antName] = struct{}{}

		ant := myAntGroup.AntsDb[antName]
		ant.CurrentRoomName = 0
		ant.StepNumber = 0
		ant.VisitedRoomsArr = []mt{0}
		myAntGroup.AntsDb[antName] = ant
	}

}
