package util

import (
	"fmt"
)

func (myGraph Graph)InitGraph(tunnelArr []string) error {

	myGraph.startRoomName = 0
	myGraph.endRoomName = 6

	myGraph.rooms = make(map[uint]Room)
	
	err := myGraph.setTunnelToGraph(tunnelArr)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	return nil
}
