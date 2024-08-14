package modelpk

func (theModel *Model) ResetFactory() {
	theModel.Graph.Reset()
	theModel.AntGroup.Reset()
}
