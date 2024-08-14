package graphpk

import (
	"fmt"
)


// ---------------------------------
func (myGraph *Graph) Print() {
	fmt.Println("\n----------Graph------------")
	str, _ := myGraph.ToString()
	fmt.Println(str)
	fmt.Println("---------------------------")
}

//------------------------------------