package data

import (
	"fmt"
)

func (myGraph *Graph) PrintGraph() {
	l := len(myGraph.Rooms)
	fmt.Println("Number of rooms:", l)
	fmt.Println("Start Room:", myGraph.StartRoomName)
	fmt.Println("End Room:", myGraph.EndRoomName)

	fmt.Println("-------Rooms:------")
	for i := range myGraph.Rooms {
		room := myGraph.Rooms[i]
		fmt.Printf("Room %v: Tunnels: %v, EmptySeats: %d, MaxSeats: %d\n", room.name, room.Tunnels, room.emptySeats, room.maxSeats)

	}
	//fmt.Println("Graph:",myGraph)
	//fmt.Println("len:",myGraph)

}
