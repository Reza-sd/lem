package graphpk

import (
	"fmt"
)

// ---------------------------------
func (myGraph *Graph) Print() {
	fmt.Println("\n----------Graph------------")
	str := myGraph.ToString()
	fmt.Println(str)
	fmt.Println("---------------------------")
}

//------------------------------------
