package antpk

import (
	"fmt"
)

func (allAnts *AntGroup) PrintAllAnts() {
	fmt.Println("-----------myAnts-------------")
	for i := 1; i <= allAnts.NumberOfAnts; i++ {
		fmt.Println(allAnts.AntsMap[i])
	}

	fmt.Println("------------------------")
}
