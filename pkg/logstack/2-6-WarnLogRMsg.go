package logstack

// ---------------WarnLogRErrMsg------------------
func (l *Logger) WarnLogRErrMsg(FuncName string, OperationName string, err any, operationDescription string) error {

	l.WarnLog(FuncName, OperationName, operationDescription, err)
	return errGenerator(FuncName, OperationName, err)
}

//------------------------------------
