package main

import (
	"main/internal/simulationpk"
)

// "fmt"
// "main/internal/antpk"
// graphpk "main/internal/graphpk"


type Lem struct{
	NumberOfAnts int
	StartRoom string
	EndRoom string
	TunnelArr []string

}

func main() {


	lem1 := simulationpk.Lem {
		NumberOfAnts: 3,
		StartRoom: "1",
		EndRoom: "0",
		TunnelArr : []string{"0-4", "0-6", "1-3", "4-3", "5-2", "3-5", "4-2", "2-1", "7-6", "7-2", "7-4", "6-5"},
	}

	lem1.Run(200)

}
