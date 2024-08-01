package antpk

import (
	"fmt"
)

func (allAnts *AntGroup) ToString() string {
	//fmt.Println("-----------myAnts-------------")
	var antGroupString string

	for i := 1; i <= allAnts.NumberOfAnts; i++ {
		fmt.Println(allAnts.AntsMap[i])
		ant := fmt.Sprintf("%v \n", allAnts.AntsMap[i])
		antGroupString += ant
	}
	return antGroupString
	//fmt.Println("------------------------")
}
