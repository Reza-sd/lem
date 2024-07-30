package antpk

import (
	//"fmt"
	"main/internal/graphpk"
)

func (allAnts *Ants) MoveAllAntsOneStepRandomly(theGraph *graphpk.Graph) {

	for i := 1; i <= allAnts.AntsNumber; i++ {
		theAnt := allAnts.AntsMap[i]
		theAnt.MoveOneStepForwardRandomly(theGraph)
		//fmt.Println(theAnt)
		allAnts.AntsMap[i] = theAnt
	}
	//fmt.Println(allAnts)
	//allAnts.PrintAllAnts()
}
