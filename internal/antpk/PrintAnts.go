package antpk

import (
	"fmt"
)

func (allAnts *Ants) PrintAllAnts() {
	for i := 1; i <= allAnts.AntsNumber; i++ {
		fmt.Println(allAnts.AntsMap[i])
	}
	//fmt.Println("number of ants in end room: ", theGraph.CurrentAntsInEndRoom)
	fmt.Println()
}
