package main

import (
	"fmt"
	graphpk "main/internal/graphpk"
	"main/internal/antpk"
)

func main() {

	var myGraph graphpk.Graph
	tunnelArr := []string{"0-4", "0-6", "1-3", "4-3", "5-2", "3-5", "4-2", "2-1", "7-6", "7-2", "7-4", "6-5"}
	myGraph.InitGraph(tunnelArr, "1", "0")
	myGraph.PrintGraph()

	//move.FindAllPaths(&myGraph)

	fmt.Println("==================")
	var myAnts antpk.Ants
	myAnts.AntsInit(3, &myGraph)
	fmt.Println("myAnts.Ants= ", myAnts.AntsMap)
	fmt.Println("-----------------")
	myAnts.MoveAllAntsOneStepRandomly(&myGraph)
	myAnts.MoveAllAntsOneStepRandomly(&myGraph)
	myAnts.MoveAllAntsOneStepRandomly(&myGraph)
	myAnts.MoveAllAntsOneStepRandomly(&myGraph)
	myAnts.MoveAllAntsOneStepRandomly(&myGraph)
	myAnts.MoveAllAntsOneStepRandomly(&myGraph)
	//myAnts.Move(&myGraph)
	// ant1 := myAnts.AntsMap[1]
	// ant2 := myAnts.AntsMap[2]
	// ant3 := myAnts.AntsMap[3]
	// fmt.Println("\n step 0:")
	// fmt.Println(ant1)
	// fmt.Println(ant2)
	// fmt.Println(ant3)
	// fmt.Println("\n step 1:")
	// ant1.MoveOneStepForwardRandomly(&myGraph)
	// ant2.MoveOneStepForwardRandomly(&myGraph)
	// ant3.MoveOneStepForwardRandomly(&myGraph)
	// fmt.Println("\n step 2:")
	// ant1.MoveOneStepForwardRandomly(&myGraph)
	// ant2.MoveOneStepForwardRandomly(&myGraph)
	// ant3.MoveOneStepForwardRandomly(&myGraph)
	// fmt.Println("\n step 3:")
	// ant1.MoveOneStepForwardRandomly(&myGraph)
	// ant2.MoveOneStepForwardRandomly(&myGraph)
	// ant3.MoveOneStepForwardRandomly(&myGraph)
	// fmt.Println("\n step 4:")
	// ant1.MoveOneStepForwardRandomly(&myGraph)
	// ant2.MoveOneStepForwardRandomly(&myGraph)
	// ant3.MoveOneStepForwardRandomly(&myGraph)

}
