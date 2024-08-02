package graphpk

import (
	"fmt"
)

func (myGraph *Graph) setTunnelToGraph(tunnelArr []string) error {

	for _, tunnel := range tunnelArr {
		a, b, err := splitStringToUint(tunnel)
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}
		// Retrieve the Room from the map
		roomA := myGraph.Rooms[a]
		roomB := myGraph.Rooms[b]

		roomA.Name = a
		roomB.Name = b
		// Modify the Room's tunnels
		roomA.Tunnels = append(roomA.Tunnels, b)
		roomB.Tunnels = append(roomB.Tunnels, a)
		// Put the modified Room back into the map
		if a == myGraph.StartRoomName || a == myGraph.EndRoomName {
			roomA.MaxSeats = 100000

		} else {
			roomA.MaxSeats = 1

		}

		if b == myGraph.StartRoomName || b == myGraph.EndRoomName {
			roomB.MaxSeats = 100000

		} else {
			roomB.MaxSeats = 1

		}

		roomA.EmptySeats = roomA.MaxSeats
		roomB.EmptySeats = roomB.MaxSeats

		myGraph.Rooms[a] = roomA
		myGraph.Rooms[b] = roomB

	}

	return nil
}

//--------------------------------
