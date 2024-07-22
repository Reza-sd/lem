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

	var path []string

	var allPath [][]string
	var pathStr string

	var allPathStr []string

	for i := 0; i < 500; i++ {
		pathStr = ""
		path = nil
		path = append(path, myGraph.StartRoomName)
		currentRoomName := myGraph.StartRoomName
		pathStr = currentRoomName

		numberOfTry := 0

		for {
			numberOfTry++
			if currentRoomName == myGraph.EndRoomName || numberOfTry > 1000 {
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
			pathStr += "-" + currentRoomName

		}
		//fmt.Println(path)
		if path[len(path)-1]!= myGraph.EndRoomName{
			continue
		}
		//fmt.Println(pathStr)
		if slices.Contains(allPathStr, pathStr){
			continue
		}


		allPathStr = append(allPathStr, pathStr)
		allPath = append(allPath, path)
	}
	fmt.Println("-------All paths:-----------")
	fmt.Println(len(allPath))
	for _, path := range allPath {
		
		fmt.Println(path)
		//fmt.Println(allPathStr)
	}
	fmt.Println(len(allPathStr))
	for _, pathStr := range allPathStr {
		fmt.Println(pathStr)
	}
	//----------------------------------------------------------------
	//var allPathNo [][]string

	// --------------------------------
}
