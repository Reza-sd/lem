package graphpk

func GraphCopy(baseGraph Graph, secondGraph *Graph) {

	secondGraph.StartRoomName= baseGraph.StartRoomName
	secondGraph.EndRoomName= baseGraph.EndRoomName
	//secondGraph.CurrentAntsInEndRoom = 0
	secondGraph.Rooms = make(map[string]Room)

	for roomName, room := range baseGraph.Rooms {
        secondGraph.Rooms[roomName] = room
    }
}