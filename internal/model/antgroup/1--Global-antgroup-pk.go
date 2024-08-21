package antgroup

import (
	antpk "main/internal/model/ant"
	logstack "main/pkg/logstack"
)

// ===========data type=================
type Mtag = uint16 //my type ant group pk
type Ant = antpk.Ant

//============const==============================

const (
	//MaxHandleableAntsNumber = 1100
	pkgName = "antgroup"
	//LogFilesDirectory =
	rootFromAntpk           = "../.."
	LogFilesDirectory       = rootFromAntpk + "/logs/"
	MaxHandleableAntsNumber = Mtag(1100)
)

//======================================================

var (
	logger = logstack.LogCollector{
		PackageName: pkgName,
		LogToFile:   true,
		LogToCli:    true,
	}
)

//======================================================
