package logstack

import (
	"fmt"
	"log/slog"
)

//import "go/types"

// ----------------------------------------
func LogMsg(packageName string, FuncName string, OperationName string, description string, RetunedError interface{}) (string, []slog.Attr) {

	//errString := fmt.Sprintf("%v", RetunedError)

	return description, []slog.Attr{
		slog.String("package", packageName),
		slog.String("func", FuncName),
		slog.String("op", OperationName),
		slog.String("error", fmt.Sprintf("%v", RetunedError)),
	}

}

//----------------------------------------------
