package util

import (
	"fmt"
)

func ReadTunnel(tunnelArr []string) error {

	fmt.Println("ReadTunnel")

	//tunnelArr := []string{"0-4", "0-6","1-3","4-3"}
	fmt.Println(tunnelArr[0])

	myGraph := Graph{
		rooms: make(map[uint]Room),
	}

	myGraph.startRoomName = 0
	myGraph.endRoomName = 6

	for _, tunnel := range tunnelArr {
		tunnelFrom, tunnelTo, err := splitStringToUint(tunnel)
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}
		// Retrieve the Room from the map
		room := myGraph.rooms[tunnelFrom]
		room.name = tunnelFrom
		// Modify the Room's tunnels
		room.tunnels = append(room.tunnels, tunnelTo)
		// Put the modified Room back into the map
		myGraph.rooms[tunnelFrom] = room
	}

	fmt.Println(myGraph)
	return nil
}

//--------------------------------
