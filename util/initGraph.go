package util

import (
	"fmt"
)

func (myGraph Graph)InitGraph(tunnelArr []string,start, end uint) error {

	myGraph.startRoomName = start
	myGraph.endRoomName = end

	myGraph.rooms = make(map[uint]Room)

	err := myGraph.setTunnelToGraph(tunnelArr)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	return nil
}
