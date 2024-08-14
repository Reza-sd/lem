package modelpk

import (
	
	antpk "main/internal/model/antpk"
	graphpk "main/internal/model/graphpk"
)
//-------------------------------------------------
type z = uint16

type Lem struct {
	NumberOfAnts z
	StartRoom    z
	EndRoom      z
	TunnelArr    []z

}

type Model struct {
	Graph    graphpk.Graph
	AntGroup antpk.AntGroup
}
//-------------------------------------------------