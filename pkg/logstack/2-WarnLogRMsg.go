package logstack

// ---------------WarnLogRMsg------------------
func (l *LogCollector) WarnLogRMsg(FuncName string, OperationName string, err interface{}, operationDescription string) error {

	l.WarnLog(FuncName, OperationName, err, operationDescription)
	return errMsg(FuncName, OperationName, err)
}

//------------------------------------
