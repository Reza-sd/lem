package modelpk

import (
	antgrouppk "main/internal/model/antgroup"
	graphpk "main/internal/model/graph"
)

// -------------------------------------------------
type Mtm = antgrouppk.Mtag //my type Model
type GraphType = graphpk.Graph
type AntGroupType = antgrouppk.AntGroup


type Lem struct {
	NumberOfAnts Mtm
	EndRoom      Mtm
	TunnelMap    map[Mtm][]Mtm
}

type Model struct {
	Graph    GraphType
	AntGroup AntGroupType
}

var (
	mySampleGraph =graphpk.SampleGragh{}
	mySampleAntGroup = antgrouppk.SampleAntGroup{}
)
//-------------------------------------------------
