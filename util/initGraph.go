package util

import (
	"fmt"
)

func (myGraph *Graph) InitGraph(tunnelArr []string, start, end uint) error {
	// set start and end room
	myGraph.startRoomName = start
	myGraph.endRoomName = end
	// Initialize rooms map
	myGraph.rooms = make(map[uint]Room)

	err := myGraph.setTunnelToGraph(tunnelArr)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	//fmt.Println(myGraph)
	return nil
}
