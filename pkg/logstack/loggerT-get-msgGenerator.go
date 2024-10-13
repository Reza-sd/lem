package logstack

import (
	"fmt"
	"log/slog"
)

// ----------------------------------------
func (get *getter[T]) msgGenerator(errCode T, desArr []any) (string, slog.Attr) {

	desStr := desArrToString(desArr)
	// we can separate method.func.,... by string dot separator
	return fmt.Sprintf("%v", get.desForErrCode(errCode)), slog.Group("",
		slog.String("pk", get.pkgName()),
		slog.String("des", desStr),
		//slog.String("func", FuncName),
		//slog.String("op", OperationName),
		//slog.String("errMsg", fmt.Sprintf("%v", errMsg)),
	)

}

// ----------------------------------------------
func desArrToString(desArr []any) string {
	var str string
	for _, item := range desArr {
		str = str + fmt.Sprintf("%v", item)
	}
	return str
}
