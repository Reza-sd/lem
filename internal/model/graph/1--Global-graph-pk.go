package graphpk

import (
	logstack "main/pkg/logstack"
)

// =========================================================
type Mtg = uint16

const (

	//MaxHandleableAntsNumber = 200
	pkgName = "graphpk"
	//LogFilesDirectory =
	//rootFromAntpk     = "../.."
	//LogFilesDirectory = rootFromAntpk + "/logs/"
)

// ------------------------------------------
var (
	logger = logstack.LogCollector{
		PackageName: pkgName,
	}
)

//=========================================================
