package move

import (
	"fmt"
	"main/internal/data"
)

func (allAnts *Ants) MoveAllAntsOneStepRandomly(theGraph *data.Graph) {

	for i:=1;i<=allAnts.AntsNumber;i++{
		theAnt :=allAnts.AntsMap[i]
		theAnt.MoveOneStepForwardRandomly(theGraph)
		//fmt.Println(theAnt)
		allAnts.AntsMap[i]=theAnt
	}
	fmt.Println(allAnts)
}