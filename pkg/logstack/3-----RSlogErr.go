package logstack

import (
	"log/slog"
)

// -------------------RSlogErr-------------------------
func RSlogErr(packageName, funcName, opName, opDes string, err any) error {
	msg, agrs := LogMsgGenerator(packageName, funcName, opName, opDes, err)

	slog.Error(msg, agrs)
	return errGenerator(funcName, opName, err)
}

//----------------------------------------------------
