package graphpk

import (
//"fmt"
)

func (myGraph *Graph) Reset() {
	myGraph.CurrentAntsInEndRoom = 0

	for _, room := range myGraph.Rooms {

		room.EmptySeats = 1
		room.maxSeats = 1

		if room.name == myGraph.StartRoomName || room.name == myGraph.EndRoomName {
			room.maxSeats = 100000
			room.EmptySeats = room.maxSeats
		}
	}

}
