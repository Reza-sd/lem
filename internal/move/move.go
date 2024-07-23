package move

import (
	"fmt"
	data "main/internal/data"
	//"math/rand"
)

func (myAnts *Ants) Move(myGraph *data.Graph) {

	fmt.Println("Move Ants")

	for _, ant := range myAnts.Ants {
		fmt.Println("Ant:", ant.Name, "Current Room:", ant.CurrentRoomName)
		//fmt.Println("Shortest Path:", ant.ShortestPath)
	}
	//------------------------------

	// what are the next free rooms in the graph
	currentRoomName := myAnts.Ants[1].CurrentRoomName
	fmt.Println(currentRoomName)

	// nextRooms := myGraph.GetFreeRooms(currentRoom)
	// for w := 0; w < 100; w++ {
	// 	if myGraph.Paths[w] != nil {
	// 		fmt.Println("length: ", w, "paths: ", myGraph.Paths[w])
	// 		break
	// 	}
	// }
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
	myAnts.ChooseTheBestPath(myGraph, currentRoomName)

	// ------------------------------
}

// ==================
func (myAnts *Ants) ChooseTheBestPath(myGraph *data.Graph, currentRoomName string) {
	currentRoom := myGraph.Rooms[currentRoomName]
	currentTunnerArr := currentRoom.Tunnels
	//lengthCurrentTunnerArr := len(currentTunnerArr)
	fmt.Println(currentTunnerArr)
	for _, tunnelto := range currentTunnerArr {
		fmt.Println(tunnelto)
	}
	//-----------------
	// -----simulate not emtpy--------
	roomTest := myGraph.Rooms["6"]
	roomTest.EmptySeats = 0
	myGraph.Rooms["6"] = roomTest
	fmt.Println(">>>>", myGraph.Rooms["6"].EmptySeats)
	//-----------
	//myGraph.StartRoomName = "7"
	Ant := myAnts.Ants[1]
	Ant.CurrentRoomName = "2"
	Ant.VisitedRooms = []string{"1", "2"}
	myAnts.Ants[1] = Ant
	//------------------------
	fmt.Println("-->>", AntNextBestAvaliableRoom(myGraph, myAnts, 1))
	NextBestAvaliableRoom := AntNextBestAvaliableRoom(myGraph, myAnts, 1)
	if NextBestAvaliableRoom == "" {
		fmt.Println("No Path Found")
		return
	}
	fmt.Println("Next Best Avaliable Room= ", NextBestAvaliableRoom)
	//myAnts.Ants[1].currentRoomName = NextBestAvaliableRoom

	//-------
}

//==================
