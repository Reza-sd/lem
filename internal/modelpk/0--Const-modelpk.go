package modelpk

import (
	//"fmt"
	"main/internal/modelpk/antpk"
	graphpk "main/internal/modelpk/graphpk"
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
	BaseGraph graphpk.Graph
	BaseAnts  antpk.AntGroup
}
