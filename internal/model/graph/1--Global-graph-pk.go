package graphpk

import (
	logstack "main/pkg/logstack"
)

// =========================================================
type Mtg = uint16
//type int = uint8

const (

	//MaxHandleableAntsNumber = 200
	pkgName   = "graphpk"
	logToFile = false
	logToCli  = false
	//LogFilesDirectory =
	//rootFromAntpk     = "../.."
	//LogFilesDirectory = rootFromAntpk + "/logs/"
)

// ------------------------------------------
var (
	logger = logstack.LogCollector{
		PackageName: pkgName,
		LogToFile:   logToFile,
		LogToCli:    logToCli,
	}
)

//=========================================================
