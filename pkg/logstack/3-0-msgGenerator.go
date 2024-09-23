package logstack

import (
	"fmt"
	"log/slog"
)

// ----------------------------------------
func (l *LoggerT) msgGenerator(errCode errT) (string, slog.Attr) {
	// we can separate method.func.,... by string dot separator
	return fmt.Sprintf("%v", l.get.DesForErrCode(errCode)), slog.Group("",
		slog.String("pk", l.get.pkgName()),
		//slog.String("func", FuncName),
		//slog.String("op", OperationName),
		//slog.String("errMsg", fmt.Sprintf("%v", errMsg)),
	)

}

//----------------------------------------------
