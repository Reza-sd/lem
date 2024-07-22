package move

import (
	"fmt"
	data "main/internal/data"
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

	// for {

	// }

	fmt.Println("Finding")
	fmt.Println(myGraph)

}
