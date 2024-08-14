package modelpk

import (
	antpk "main/internal/model/antpk"
	graphpk "main/internal/model/graphpk"
)

// -------------------------------------------------
type Mtm = uint16 //my type Model

type Lem struct {
	NumberOfAnts Mtm
	EndRoom      Mtm
	TunnelMap    map[Mtm][]Mtm
}

type Model struct {
	Graph    graphpk.Graph
	AntGroup antpk.AntGroup
}

//-------------------------------------------------
