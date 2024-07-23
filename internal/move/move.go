package move

import (
	"fmt"
	data "main/internal/data"
	//"math/rand"
)

func (myAnts *Ants) Move(myGraph *data.Graph) {

	fmt.Println("Move Ants")

	for _, ant := range myAnts.Ants {
		fmt.Println("Ant:", ant.Name, "Current Room:", ant.currentRoomName)
		//fmt.Println("Shortest Path:", ant.ShortestPath)
	}
	//------------------------------

	// what are the next free rooms in the graph
	currentRoomName := myAnts.Ants[1].currentRoomName
	fmt.Println(currentRoomName)

	// nextRooms := myGraph.GetFreeRooms(currentRoom)
	for w := 0; w < 100; w++ {
		if myGraph.Paths[w] != nil {
			fmt.Println("length: ", w, "paths: ", myGraph.Paths[w])
			break
		}
	}
	/*
	   currentroom
	   nextrooms?
	   if nextroomsfree?
	   nextfreerooms?
	   currentroom-nextfreeroom in whitch path?
	   path with lower weight will select
	*/
	// whats the random next room
	// currentRoom := myGraph.Rooms[currentRoomName]
	// currentTunnerArr := currentRoom.Tunnels
	// lengthCurrentTunnerArr := len(currentTunnerArr)

	// randomNextRoomIndex := rand.Intn(lengthCurrentTunnerArr)
	// nextRandomRoomName := currentTunnerArr[randomNextRoomIndex]
	// fmt.Println("nextRandomRoomName=", nextRandomRoomName)
	chooseTheBestPath(myGraph, currentRoomName)


	// ------------------------------
}
//==================
func chooseTheBestPath(myGraph *data.Graph, currentRoomName string){
	currentRoom := myGraph.Rooms[currentRoomName]
	currentTunnerArr := currentRoom.Tunnels
	//lengthCurrentTunnerArr := len(currentTunnerArr)
	fmt.Println(currentTunnerArr)
	for _,tunnelto :=range currentTunnerArr {
		fmt.Println(tunnelto)
	}
}

//==================