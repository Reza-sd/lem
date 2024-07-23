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
	//-------test----
	//myGraph.StartRoomName = "7"
	// Ant1 := myAnts.Ants[1]
	// Ant1.CurrentRoomName = "2"
	// Ant1.VisitedRooms = []string{"1", "2"}
	// myAnts.Ants[1] = Ant1
	// //----------test--------------
	// fmt.Println("-->>", AntNextBestAvaliableRoom(myGraph, myAnts, 1))
	// NextBestAvaliableRoom1 := AntNextBestAvaliableRoom(myGraph, myAnts, 1)
	// if NextBestAvaliableRoom1 == "" {
	// 	fmt.Println("No Path Found")
	// 	return
	// }
	// fmt.Println("Next Best Avaliable Room= ", NextBestAvaliableRoom1)
	//myAnts.Ants[1].currentRoomName = NextBestAvaliableRoom
	//--------------------------------------------
	for i:=1;i<=len(myAnts.Ants);i++ {
		Ant := myAnts.Ants[i]
		NextBestAvaliableRoom := AntNextBestAvaliableRoom(myGraph, myAnts, i)
		if NextBestAvaliableRoom == "" {
            fmt.Println("No Path Found for Ant: ", Ant.Name)
            continue
        }
		
		//myGraph.Rooms[NextBestAvaliableRoom].EmptySeats
		currentRoom := myGraph.Rooms[Ant.CurrentRoomName]
		newNextRoom := myGraph.Rooms[NextBestAvaliableRoom]

		currentRoom.EmptySeats=1 //free seat
		newNextRoom.EmptySeats=0 //not free seat

		myGraph.Rooms[Ant.CurrentRoomName] = currentRoom
		myGraph.Rooms[NextBestAvaliableRoom] = newNextRoom

		
		
		Ant.CurrentRoomName = NextBestAvaliableRoom
		Ant.VisitedRooms = append(Ant.VisitedRooms, Ant.CurrentRoomName)
		myAnts.Ants[i] = Ant
	}
	fmt.Println(myAnts)
	//-------
}

//==================
