package antgroup

import (
	"fmt"
)

// --------------------
func (allAnts *AntGroup) Print() {
	fmt.Println("\n----------AntGroup------------")
	str, _ := allAnts.ToString()
	fmt.Println(str)
	fmt.Println("---------------------------")
}

//-------------------------
