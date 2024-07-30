package simulationpk

import (
	//"fmt"
	"main/internal/antpk"
	graphpk "main/internal/graphpk"
)

// model + behaviour
type Lem struct {
	NumberOfAnts int
	StartRoom    string
	EndRoom      string
	TunnelArr    []string

	baseModel Model
}

type Model struct {
	baseGraph graphpk.Graph
	baseAnts  antpk.Ants
}
