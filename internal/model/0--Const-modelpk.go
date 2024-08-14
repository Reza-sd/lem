package modelpk

import (
	//"fmt"
	antpk "main/internal/model/antpk"
	graphpk "main/internal/model/graphpk"
)

type z = uint16

// model + behaviour
type Lem struct {
	NumberOfAnts z
	StartRoom    z
	EndRoom      z
	TunnelArr    []z

	//baseModel Model
}

type Model struct {
	Graph    graphpk.Graph
	AntGroup antpk.AntGroup
}
