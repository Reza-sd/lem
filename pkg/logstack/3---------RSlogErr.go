package logstack

import (
	"log/slog"
)

// -------------------RSlogErr-------------------------
func RSlogErr(packageName, funcName, opName, opDes string, err error) error {
	s := LogMsg(packageName, funcName, opName, opDes, err)

	slog.Error(s[0], s[1:])
	return errGenerator(funcName, opName, err)
}

//----------------------------------------------------
