package util

import (
	"fmt"
)

func InitGraph(tunnelArr []string) error {

	myGraph.startRoomName = 0
	myGraph.endRoomName = 6

	err := setTunnelToGraph(tunnelArr)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	return nil
}
