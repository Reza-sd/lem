package logstack

// ---------------------------------
func (l *LogCollector) ErrLogRMsg(FuncName string, OperationName string, err error, operationDescription string) error {

	l.ErrLog(FuncName, OperationName, err, operationDescription)

	return errMsg(FuncName, OperationName, err) //optional return
}

// ---------------------------------
