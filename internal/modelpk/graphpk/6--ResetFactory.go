package graphpk

func (myGraph *Graph) ResetFactory() {

	// for roomName, _ := range theGraph.Rooms {

	// 	roomObject := theGraph.Rooms[roomName]

	// 	if roomName != theGraph.StartRoomName && roomName != theGraph.EndRoomName {
	// 		roomObject.EmptySeats = 1
	// 		theGraph.Rooms[roomName] = roomObject
	// 	}

	// }
	for key := range myGraph.RoomAvailableDb {
		myGraph.RoomAvailableDb[key] = true
	}
}
