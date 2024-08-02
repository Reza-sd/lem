package graphpk

import (
	"fmt"
)

func (myGraph *Graph) PrintGraph() {
	fmt.Println("=============myGraph:==============")
	l := len(myGraph.Rooms)
	fmt.Println("Number of rooms:", l)
	fmt.Println("Start Room:", myGraph.StartRoomName)
	fmt.Println("End Room:", myGraph.EndRoomName)

	fmt.Println("-------Rooms:------")

	for i := range myGraph.Rooms {
		room := myGraph.Rooms[i]
		fmt.Printf("Room %v: Tunnels: %v, EmptySeats: %d, MaxSeats: %d\n", room.Name, room.Tunnels, room.EmptySeats, room.MaxSeats)

	}
	//fmt.Println("Graph:",myGraph)
	//fmt.Println("len:",myGraph)
	fmt.Println("======================================")

}
