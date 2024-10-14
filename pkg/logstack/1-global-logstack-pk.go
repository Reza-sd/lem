package logstack

import (
	"log/slog"
	"os"
	"time"
)

// --------------const-----------------------------

type errType interface {
	uint8 | uint16 | uint32
}

const (
	_PKG_NAME         = "logStack"
	_LOG_FILE_NAME_PREFIX = "z-log-"
	_LOG_FILES_DIRECTORY = "./"
	OP = 2
)

// -----------var----------------------
var (
	//----------------------------------

	logHandlerOptsCli = &slog.HandlerOptions{
		Level: slog.LevelDebug,
		//AddSource: true,
	}
	logHandler  = slog.NewTextHandler(os.Stderr, logHandlerOptsCli)
	loggerToCli = slog.New(logHandler)

	//--------------------------
	todayDate      = time.Now().Format("2006-01-02")
	logFileAddress = _LOG_FILES_DIRECTORY + _LOG_FILE_NAME_PREFIX + todayDate + ".json"
	logFile, _     = os.OpenFile(logFileAddress, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	//-------------------------------
	logHandlerOptsFile = &slog.HandlerOptions{
		Level: slog.LevelDebug,
		//AddSource: true,
	}
	loggerToFile = slog.New(slog.NewJSONHandler(logFile, logHandlerOptsFile))
	//--------------------------
)
