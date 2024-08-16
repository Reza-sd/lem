package logstack

// ---------------WarnLogRErrMsg------------------
func (l *LogCollector) WarnLogRErrMsg(FuncName string, OperationName string, err error, operationDescription string) error {

	l.WarnLog(FuncName, OperationName, err, operationDescription)
	return errMsg(FuncName, OperationName, err)
}

//------------------------------------
