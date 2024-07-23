package data

import (
	"fmt"
)

func (myGraph *Graph) InitGraph(tunnelArr []string, start, end string) error {
	// set start and end room
	myGraph.StartRoomName = start
	myGraph.EndRoomName = end
	// Initialize rooms map
	myGraph.Rooms = make(map[string]Room)

	err := myGraph.setTunnelToGraph(tunnelArr)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	return nil
}
