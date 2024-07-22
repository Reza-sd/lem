package util

import (
	"fmt"
)

func (myGraph *Graph) PrintGraph() {
l := len(myGraph.rooms)
fmt.Println("Number of rooms:", l)
fmt.Println("Start Room:", myGraph.startRoomName)
fmt.Println("End Room:", myGraph.endRoomName)

    fmt.Println("-------Rooms:------")
for i := range myGraph.rooms{
	room := myGraph.rooms[i]
    fmt.Printf("Room %d: Tunnels: %v, EmptySeats: %d, MaxSeats: %d\n", room.name, room.tunnels, room.emptySeats, room.maxSeats)
    
}
	//fmt.Println("Graph:",myGraph)
	//fmt.Println("len:",myGraph)

}