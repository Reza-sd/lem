package logstack

import (
	//"errors"
	"log/slog"
	"os"
	"time"
)

// --------------const-----------------------------
type errT = uint8

//type errArrT = []errT

const (
	pkgName           = "logStack"
	logFileNamePrefix = "z-log-"
	LogFilesDirectory = "./"
)

// -----------var----------------------
var (
	//-----------------------------------
	//----------------------------------

	logHandlerOptsCli = &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}
	logHandler  = slog.NewTextHandler(os.Stderr, logHandlerOptsCli)
	loggerToCli = slog.New(logHandler)

	//--------------------------
	todayDate      = time.Now().Format("2006-01-02")
	logFileAddress = LogFilesDirectory + logFileNamePrefix + todayDate + ".json"
	logFile, _     = os.OpenFile(logFileAddress, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	//-------------------------------
	logHandlerOptsFile = &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}
	loggerToFile = slog.New(slog.NewJSONHandler(logFile, logHandlerOptsFile))
	//--------------------------
)
