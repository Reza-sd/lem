package move

import (
	"fmt"
	data "main/internal/data"
	//"math/rand"
)

func (Ant *Ant) moveOneStepForward (myGraph *data.Graph) {
fmt.Println("move theAnt")

//-----check if its already in End room----
if Ant.CurrentRoomName == myGraph.EndRoomName {
	return
}
//nextRandomRoomName := nextRandomRoomName ()

}