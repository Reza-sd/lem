package logstack

import (
	"errors"
	"log/slog"
	"os"
	"time"
)

// --------------const-----------------------------

const (
	pkgName = "logStack"
	filePrefix = "z-log-"
)

// -----------var----------------------
var (
	//loggerToFile *slog.Logger
	//loggerToCli  *slog.Logger

	//logger LogCollector
	//----------------------------
	logHandlerOptsCli = &slog.HandlerOptions{
		Level: slog.LevelDebug,
		//AddSource: true,
	}
	logHandler        = slog.NewTextHandler(os.Stderr, logHandlerOptsCli)
	loggerToCli       = slog.New(logHandler)
	LogFilesDirectory = "./"
	//--------------------------
	todayDate      = time.Now().Format("2006-01-02")
	logFileAddress = LogFilesDirectory + filePrefix + todayDate + ".json"
	logFile, _     = os.OpenFile(logFileAddress, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	//-------------------------------
	logHandlerOptsFile = &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}
	loggerToFile = slog.New(slog.NewJSONHandler(logFile, logHandlerOptsFile))
	//--------------------------
)

// ===================struct=========================
type LogCollector struct {
	PackageName string
}

// ---------------------------------
func (l *LogCollector) ErrMsg(FuncName string, OperationName string, RetunedError error) error {
	return errMsg(FuncName, OperationName, RetunedError)
}

// ---------------------------------
func (l *LogCollector) Info(FuncName string, OperationName string, operationDescription string) {

	operationDescription = " success: " + operationDescription

	loggerToCli.Info(LogMsg(l.PackageName, FuncName, OperationName, operationDescription, nil))

	loggerToFile.Info(LogMsg(l.PackageName, FuncName, OperationName, operationDescription, nil))
}

// ---------------------------------
func (l *LogCollector) RErr(FuncName string, OperationName string, err error, operationDescription string) error {

	l.Error(FuncName, OperationName, err, operationDescription)

	return errMsg(FuncName, OperationName, err) //optional return
}

// ---------------------------------
func (l *LogCollector) Error(FuncName string, OperationName string, err error, operationDescription string) {

	operationDescription = " fail: " + operationDescription

	loggerToCli.Error(LogMsg(l.PackageName, FuncName, OperationName, operationDescription, err))
	loggerToFile.Error(LogMsg(l.PackageName, FuncName, OperationName, operationDescription, err))

}

// ---------------RWarn------------------
func (l *LogCollector) RWarn(FuncName string, OperationName string, err error, operationDescription string) error {
	l.Warn(FuncName, OperationName, err, operationDescription)
	return errMsg(FuncName, OperationName, err)
}

// ----------------RWarnStr-----------------
func (l *LogCollector) RWarnStr(FuncName string, OperationName string, errStr string, operationDescription string) error {

	err := errors.New(errStr)

	return l.RWarn(FuncName, OperationName, err, operationDescription)
}

// ----------------Warn-----------------
func (l *LogCollector) Warn(FuncName string, OperationName string, err error, operationDescription string) {

	operationDescription = " !!!: " + operationDescription

	loggerToCli.Warn(LogMsg(l.PackageName, FuncName, OperationName, operationDescription, err))
	loggerToFile.Warn(LogMsg(l.PackageName, FuncName, OperationName, operationDescription, err))

}

// -------------------RSlogErr-------------------------
func RSlogErr(packageName, funcName, opName, opDes string, err error) error {
	slog.Error(LogMsg(packageName, funcName, opName, opDes, err))
	return errMsg(funcName, opName, err)
}

//----------------------------------------------------
