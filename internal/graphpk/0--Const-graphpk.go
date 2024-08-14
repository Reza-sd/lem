package graphpk

import (
	"main/internal/logstack"
)

/*
"0-4","0-6","1-3","4-3","5-2","3-5","4-2","2-1","7-6","7-2","7-4","6-5"

L1-3 L2-2
L1-4 L2-5 L3-3
L1-0 L2-6 L3-4
L2-0 L3-0
$
*/
//==============data structure=======================
type mtg = uint16

type Graph struct {
	EndRoomName mtg

	RoomHasEmptySeatDb map[mtg]bool
	TunnelsDb          map[mtg][]mtg
}
type TunnelMap = map[mtg][]mtg
// =========================================================
const (
	//MaxHandleableAntsNumber = 200
	pkgName = "graphpk"
	//LogFilesDirectory =
	//rootFromAntpk     = "../.."
	//LogFilesDirectory = rootFromAntpk + "/logs/"
)

var (
	logger = logstack.LogCollector{
		PackageName: pkgName,
	}
)

//=========================================================
