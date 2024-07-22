package main

import (
	"fmt"
	data "main/internal/data"
	move "main/internal/move"
)

func main() {
	//var tunnelArr []string

	var myGraph data.Graph
	//tunnelArr := []string{"0-4", "0-6", "1-3", "4-3", "5-2", "3-5", "4-2", "2-1", "7-6", "7-2", "7-4", "6-5"}
	tunnelArr := []string{"0-4", "0-6", "1-3", "4-3", "5-2", "3-5", "4-2", "2-1", "7-6", "7-2", "7-4", "6-5"}
	myGraph.InitGraph(tunnelArr, "1", "0")
	myGraph.PrintGraph()
	//fmt.Println(myGraph)
	move.FindAllPaths(&myGraph)

	fmt.Println("==================")
	var myAnts move.Ants
	myAnts.AntsInit(3,&myGraph)
	fmt.Println(myAnts.Ants)
	fmt.Println("==================")
	// var myGraph2 data.Graph
	// tunnelArr2 := []string{"a-b", "a-c", "reza-3", "4-3","san-2"}
	// myGraph2.InitGraph(tunnelArr2, "mio", "0")
	// myGraph2.PrintGraph()
	// fmt.Println("===================")

}
