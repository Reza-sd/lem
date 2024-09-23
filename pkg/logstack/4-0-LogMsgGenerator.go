package logstack

import (
	"fmt"
	"log/slog"
)

//import "go/types"

// ----------------------------------------
func LogMsgGenerator(packageName string, FuncName string, OperationName string, description string, RetunedError any) (string, slog.Attr) {

	return description, slog.Group("",
		slog.String("pk", packageName),
		slog.String("func", FuncName),
		slog.String("op", OperationName),
		slog.String("errMsg", fmt.Sprintf("%v", RetunedError)),
	)

}

//----------------------------------------------
