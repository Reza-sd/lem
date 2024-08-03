package graphpk

import (
	"fmt"
)

func (myGraph *Graph) ToString() (string, error) {
	var graphString string

	graphString = fmt.Sprintf("\nStartRoomName= %v, EndRoomName= %v \nNumberOfAllRoom= %v, CurrentAntsInEndRoom= %v\n", myGraph.StartRoomName, myGraph.EndRoomName, myGraph.NumberOfAllRoom, myGraph.CurrentAntsInEndRoom)

	count := 0
	for roomName := range myGraph.Rooms {
		count++
		room := myGraph.Rooms[roomName]
		roomStr := fmt.Sprintf("\n%v-Room= %v, Tunnels= %v, EmptySeats= %d, MaxSeats= %d,", count, room.Name, room.Tunnels, room.EmptySeats, room.MaxSeats)
		graphString += roomStr
	}
	graphString += "\n"
	return graphString, nil
}
