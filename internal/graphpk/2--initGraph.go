package graphpk

import (
	"fmt"
)

func (myGraph *Graph) InitGraph(tunnelArr []string, start, end string) error {
	// set start and end room
	myGraph.StartRoomName = start
	myGraph.EndRoomName = end
	// Initialize rooms map
	myGraph.Rooms = make(map[string]Room)
	//myGraph.UsedTunnelsInLastSequence = make(map[string]string)

	err := myGraph.setTunnelToGraph(tunnelArr)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	myGraph.NumberOfAllRoom = len(myGraph.Rooms)
	return nil
}
