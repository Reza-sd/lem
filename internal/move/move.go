package move

import (
	"fmt"
	data "main/internal/data"
	//"math/rand"
)

// ==================
func (myAnts *Ants) Move(myGraph *data.Graph) {
	//--------------------------------------
	for j := 1; j <= 10; j++ {
		//--------check if all Ants arrived---------------
		counter := 0
		for j := 1; j <= len(myAnts.Ants); j++ {
			if myAnts.Ants[j].CurrentRoomName == myGraph.EndRoomName {
				counter++
			}

		}
		if counter == len(myAnts.Ants) {
			break
		}

		// -------------------
		fmt.Println("---iteration=", j)

		for i := 1; i <= len(myAnts.Ants); i++ {
			Ant := myAnts.Ants[i]
			if Ant.CurrentRoomName == myGraph.EndRoomName {
				continue
			}
			NextBestAvaliableRoom := AntNextBestAvaliableRoom(myGraph, myAnts, i)
			fmt.Println("find= ", i, "next", NextBestAvaliableRoom)
			if NextBestAvaliableRoom == "" {
				//fmt.Println("No Path Found for Ant: ", Ant.Name)
				continue
			}

			//myGraph.Rooms[NextBestAvaliableRoom].EmptySeats
			currentRoom := myGraph.Rooms[Ant.CurrentRoomName]
			newNextRoom := myGraph.Rooms[NextBestAvaliableRoom]

			currentRoom.EmptySeats = 1 //free seat
			newNextRoom.EmptySeats = 0 //not free seat

			if NextBestAvaliableRoom == myGraph.EndRoomName {
				newNextRoom.EmptySeats = 10000
			}

			myGraph.Rooms[Ant.CurrentRoomName] = currentRoom
			myGraph.Rooms[NextBestAvaliableRoom] = newNextRoom

			Ant.CurrentRoomName = NextBestAvaliableRoom
			Ant.VisitedRooms = append(Ant.VisitedRooms, Ant.CurrentRoomName)
			myAnts.Ants[i] = Ant
		}
		fmt.Println(myAnts)
		//-------
	}
}

//==================
