package ant

import (
	"fmt"
)

func (myAnt *Ant) ToString()(string,error){



antString := fmt.Sprintf("antName=%v, CurrentRoomName=%v, StepNumber=%v, VisitedRooms=%v",myAnt.Name,myAnt.CurrentRoomName,myAnt.StepNumber,myAnt.VisitedRoomsArr)



	return antString , nil
}