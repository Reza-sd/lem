package modelpk

import (
// "fmt"
// "main/internal/antpk"
// graphpk "main/internal/graphpk"
)

func (theModel *Model) ResetFactory() {
	theModel.BaseGraph.ResetFactory()
	theModel.BaseAnts.ResetFactory(theModel.BaseGraph.StartRoomName)
}
