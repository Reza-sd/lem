package modelpk


func (theModel *Model) ResetFactory() {
	theModel.Graph.ResetFactory()
	theModel.AntGroup.ResetFactory()
}
