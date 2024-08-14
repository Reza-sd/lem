package modelpk

import (
// "fmt"
// "main/internal/antpk"
// graphpk "main/internal/graphpk"
)

func (theModel *Model) ResetFactory() {
	theModel.Graph.ResetFactory()
	theModel.AntGroup.ResetFactory()
}
