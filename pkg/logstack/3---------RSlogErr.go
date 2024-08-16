package logstack

import (
	"log/slog"
)

// -------------------RSlogErr-------------------------
func RSlogErr(packageName, funcName, opName, opDes string, err error) error {
	slog.Error(LogMsg(packageName, funcName, opName, opDes, err))
	return errMsg(funcName, opName, err)
}

//----------------------------------------------------
