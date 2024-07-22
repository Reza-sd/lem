package move

import (
	"fmt"
	data "main/internal/data"
)

func (myAnts *Ants)Move (myGraph *data.Graph){

	fmt.Println("Move Ants")

	for _, ant := range myAnts.Ants {
        fmt.Println("Ant:", ant.Name, "Current Room:", ant.CurrentRoom)
        //fmt.Println("Shortest Path:", ant.ShortestPath)
    }
//------------------------------

//what are the next free rooms in the graph
//currentRoom := myAnts.Ants[1].CurrentRoom
//nextRooms := myGraph.GetFreeRooms(currentRoom)



//------------------------------
}