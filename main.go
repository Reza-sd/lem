package main

import (
	"fmt"
	"main/util"
)

func main() {
	//var tunnelArr []string

	var myGraph util.Graph
	tunnelArr := []string{"0-4", "0-6", "1-3", "4-3"}
	myGraph.InitGraph(tunnelArr, 0, 6)

	fmt.Println(myGraph)

}
