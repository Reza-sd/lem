package util

import (
	"fmt"
)

func (myGraph Graph) setTunnelToGraph(tunnelArr []string) error {

	//fmt.Println("ReadTunnel")

	//tunnelArr := []string{"0-4", "0-6","1-3","4-3"}
	//fmt.Println(tunnelArr[0])

	for _, tunnel := range tunnelArr {
		a, b, err := splitStringToUint(tunnel)
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}
		// Retrieve the Room from the map
		roomA := myGraph.rooms[a]
		roomB := myGraph.rooms[b]

		roomA.name = a
		roomB.name = b
		// Modify the Room's tunnels
		roomA.tunnels = append(roomA.tunnels, b)
		roomB.tunnels = append(roomB.tunnels, a)
		// Put the modified Room back into the map
		if a == myGraph.startRoomName || a == myGraph.endRoomName {
			roomA.maxSeats = 100000

		} else {
			roomA.maxSeats = 1

		}

		if b == myGraph.startRoomName || b == myGraph.endRoomName {
			roomB.maxSeats = 100000

		} else {
			roomB.maxSeats = 1

		}

		roomA.emptySeats = roomA.maxSeats
		roomB.emptySeats = roomB.maxSeats

		myGraph.rooms[a] = roomA
		myGraph.rooms[b] = roomB

	}

	//fmt.Println(myGraph)
	return nil
}

//--------------------------------
