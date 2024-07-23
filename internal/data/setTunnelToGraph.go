package data

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
		roomA := myGraph.Rooms[a]
		roomB := myGraph.Rooms[b]

		roomA.name = a
		roomB.name = b
		// Modify the Room's tunnels
		roomA.Tunnels = append(roomA.Tunnels, b)
		roomB.Tunnels = append(roomB.Tunnels, a)
		// Put the modified Room back into the map
		if a == myGraph.StartRoomName || a == myGraph.EndRoomName {
			roomA.maxSeats = 100000

		} else {
			roomA.maxSeats = 1

		}

		if b == myGraph.StartRoomName || b == myGraph.EndRoomName {
			roomB.maxSeats = 100000

		} else {
			roomB.maxSeats = 1

		}

		roomA.EmptySeats = roomA.maxSeats
		roomB.EmptySeats = roomB.maxSeats

		myGraph.Rooms[a] = roomA
		myGraph.Rooms[b] = roomB

	}
	// -------------
	roomTest := myGraph.Rooms["7"]
	roomTest.EmptySeats = 0
	myGraph.Rooms["7"] = roomTest
	fmt.Println(">>>>", myGraph.Rooms["7"].EmptySeats)
	//-----------
	//fmt.Println(myGraph)
	return nil
}

//--------------------------------
