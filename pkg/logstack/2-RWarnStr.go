package logstack

import (
	"errors"
)

// ----------------RWarnStr-----------------
func (l *LogCollector) RWarnStr(FuncName string, OperationName string, errStr string, operationDescription string) error {

	err := errors.New(errStr)

	return l.RWarnErr(FuncName, OperationName, err, operationDescription)
}

//---------------------------------------------
