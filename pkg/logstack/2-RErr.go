package logstack

// ---------------------------------
func (l *LogCollector) RErr(FuncName string, OperationName string, err error, operationDescription string) error {

	l.ErrorToCliFile(FuncName, OperationName, err, operationDescription)

	return errMsg(FuncName, OperationName, err) //optional return
}

// ---------------------------------
