package logstack

// ---------------RWarnErr------------------
func (l *LogCollector) RWarnErr(FuncName string, OperationName string, err error, operationDescription string) error {
	l.WarnToCliFile(FuncName, OperationName, err, operationDescription)
	return errMsg(FuncName, OperationName, err)
}

//------------------------------------
