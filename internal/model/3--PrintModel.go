package modelpk

func (theModel *Model) Print() {
	theModel.Graph.Print()
	theModel.AntGroup.Print()
}
