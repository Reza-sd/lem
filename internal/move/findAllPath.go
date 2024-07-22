package move

import (
	"fmt"
	data "main/internal/data"
	 "math/rand"
	 "slices"
)

type ant struct {
	WhereAmI string
}

func FindAllPaths(myGraph *data.Graph) {

	//var ants map[int]ant
	ants := make(map[int]ant)

	var path []string
	path = append(path, myGraph.StartRoomName)

	ants[1] = ant{WhereAmI: myGraph.StartRoomName}
	//whereIsTheNextAvailableTunnel
	//numberOfRooms := len(myGraph.Rooms)
	currentRoomName := myGraph.StartRoomName
	numberOfTry :=0


for {
	numberOfTry ++
	if currentRoomName == myGraph.EndRoomName || numberOfTry >10000 {break}

	currentRoom := myGraph.Rooms[currentRoomName]
	currentTunnerArr := currentRoom.Tunnels
	lengthCurrentTunnerArr := len(currentTunnerArr)



	randomNextRoomIndex := rand.Intn(lengthCurrentTunnerArr)
	nextRandomRoomName := currentTunnerArr[randomNextRoomIndex]

	if slices.Contains(path,nextRandomRoomName) {continue}
	
	currentRoomName = nextRandomRoomName
	path = append(path, currentRoomName)
	
}


	fmt.Println(path)
	//fmt.Println(currentRoom)
	//fmt.Println("Finding")
	//fmt.Println(myGraph)
	//fmt.Println(rand.Intn(5)) //0-4

}
