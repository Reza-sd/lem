package main

import (
	"fmt"
	data "main/internal/data"
	move "main/internal/move"
)

func main() {

	var myGraph data.Graph
	tunnelArr := []string{"0-4", "0-6", "1-3", "4-3", "5-2", "3-5", "4-2", "2-1", "7-6", "7-2", "7-4", "6-5"}
	myGraph.InitGraph(tunnelArr, "1", "0")
	myGraph.PrintGraph()

	move.FindAllPaths(&myGraph)

	fmt.Println("==================")
	var myAnts move.Ants
	myAnts.AntsInit(3, &myGraph)
	fmt.Println("myAnts.Ants= ", myAnts.AntsMap)
	fmt.Println("-----------------")
	//myAnts.Move(&myGraph)
	ant1 :=myAnts.AntsMap[1]
	fmt.Println(ant1)
	ant1.MoveOneStepForwardRandomly(&myGraph)
	fmt.Println(ant1)
}
