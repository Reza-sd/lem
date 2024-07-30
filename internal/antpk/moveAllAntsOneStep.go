package antpk

import (
	//"fmt"
	"main/internal/graphpk"
)

func (allAnts *Ants) MoveAllAntsOneStepRandomly(theGraph *graphpk.Graph) {

	for i := 1; i <= allAnts.AntsNumber; i++ {
		theAnt := allAnts.AntsMap[i]
		theAnt.MoveOneStepForwardRandomly(theGraph)

		allAnts.AntsMap[i] = theAnt
	}

}
