package logstack

import (
	"fmt"
	"log/slog"
)

// ----------------------------------------
func (l *loggerT) msgGenerator(errCode errT, des ...any) (string, slog.Attr) {

	desStr := desArrToString(des)
	// we can separate method.func.,... by string dot separator
	return fmt.Sprintf("%v", l.get.DesForErrCode(errCode)), slog.Group("",
		slog.String("pk", l.get.pkgName()),
		slog.String("des", desStr),
		//slog.String("func", FuncName),
		//slog.String("op", OperationName),
		//slog.String("errMsg", fmt.Sprintf("%v", errMsg)),
	)

}

// ----------------------------------------------
func desArrToString(des []any) string {
	var str string
	for _, item := range des {
		str = str + fmt.Sprintf("%v", item)
	}
	return str
}
