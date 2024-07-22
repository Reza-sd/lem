package move

import (
	"fmt"
	data "main/internal/data"
	"math/rand"
	"slices"
)

// type ant struct {
// 	WhereAmI string
// }

func FindAllPaths(myGraph *data.Graph) {


	var path  []string

	var allPath [][]string

	for i := 0; i < 20; i++ {
		
		path = nil
		path = append(path, myGraph.StartRoomName)
		currentRoomName := myGraph.StartRoomName
		numberOfTry := 0

		for {
			numberOfTry++
			if currentRoomName == myGraph.EndRoomName || numberOfTry > 10000 {
				break
			}

			currentRoom := myGraph.Rooms[currentRoomName]
			currentTunnerArr := currentRoom.Tunnels
			lengthCurrentTunnerArr := len(currentTunnerArr)

			randomNextRoomIndex := rand.Intn(lengthCurrentTunnerArr)
			nextRandomRoomName := currentTunnerArr[randomNextRoomIndex]

			if slices.Contains(path, nextRandomRoomName) {
				continue
			}

			currentRoomName = nextRandomRoomName
			path = append(path, currentRoomName)

		}
		fmt.Println(path)
		allPath = append(allPath,path)
	}
	fmt.Println(allPath)


}
