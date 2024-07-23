package move

import (
	"fmt"
	data "main/internal/data"
	"math/rand"
	"slices"
	//"strings"
)

// type ant struct {
// 	WhereAmI string
// }
/*
To be the first to arrive, ants will need to take the shortest path or paths. They will also need to avoid traffic jams as well as walking all over their fellow ants.

*/

func AntNextBestAvaliableRoom(myGraph *data.Graph, myAnts *Ants, AntNumber int) (nextRoom string) {

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

			Ant := myAnts.Ants[AntNumber]

			if slices.Contains(Ant.VisitedRooms, nextRandomRoomName) {
				continue
			}

			if myGraph.Rooms[nextRandomRoomName].EmptySeats == 0 {
				continue
			}

			if slices.Contains(path, nextRandomRoomName) {
				continue
			}

			currentRoomName = nextRandomRoomName
			path = append(path, currentRoomName)
			pathStr += "-" + currentRoomName

		}
		
		if path[len(path)-1] != myGraph.EndRoomName {
			continue
		}
		
		if slices.Contains(allPathStr, pathStr) {
			continue
		}

		allPathStr = append(allPathStr, pathStr)
		allPath = append(allPath, path)
	}
	fmt.Println("======All paths for ant:=====")
	
	allPathMap := make(map[int][][]string)
	

	for _, path := range allPath {
	
		allPathMap[len(path)] = append(allPathMap[len(path)], path)

	}
	fmt.Println("*** number of paths for Ant:>>", len(allPathStr))
	fmt.Println()
	
	for i := 0; i < 100; i++ {

		if allPathMap[i] == nil {
			continue
		}

		fmt.Printf("length: %d,Num %d ,paths: %v\n", i, len(allPathMap[i]), allPathMap[i])
	}
	//----------------------
	for i := 0; i < 100; i++ {

		if allPathMap[i] == nil {
			continue
		}

	
		if len(allPathMap[i]) > 0 {
			if len(allPathMap[i][0])>1	{
				return allPathMap[i][0][1]
			}
		
		}
	}
	//---------------------------------

	return ""
}
