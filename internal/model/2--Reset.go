package modelpk

func (theModel *Model) Reset() {
	theModel.Graph.Reset()
	theModel.AntGroup.Reset()
}
