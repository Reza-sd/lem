package main

import (
	"fmt"
	data "main/internal/data"
)

func main() {
	//var tunnelArr []string

	var myGraph data.Graph
	tunnelArr := []string{"0-4", "0-6", "1-3", "4-3","5-2","3-5","4-2","2-1","7-6","7-2","7-4","6-5"}
	myGraph.InitGraph(tunnelArr, 1, 0)
	myGraph.PrintGraph()
	fmt.Println("--------")

}
