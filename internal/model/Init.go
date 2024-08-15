package modelpk

func (myModel *Model) Init(myLem *Lem) error {

	myModel.AntGroup.Init(Mtm(myLem.NumberOfAnts))
	myModel.Graph.Init(myLem.TunnelMap, Mtm(myLem.EndRoom))

	return nil
}
