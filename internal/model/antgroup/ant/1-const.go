package ant

import (
	//"main/internal/logstack"
)
//=======data structure========================
type Mta = uint16 //my type ant pk

type Ant struct {
	Name            Mta
	CurrentRoomName Mta
	StepNumber      Mta
	VisitedRoomsArr []Mta
}
//--------------------------------
const (
	
	//pkgName = "ant"

)

var (
	// logger = logstack.LogCollector{
	// 	PackageName: pkgName,
	// }
)
//--------------------------------
