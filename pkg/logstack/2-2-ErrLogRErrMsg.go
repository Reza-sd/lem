package logstack

// ---------------------------------
func (l *LogCollector) ErrLogRErrMsg(FuncName string, OperationName string, operationDescription string, err interface{}) error {

	l.ErrLog(FuncName, OperationName, operationDescription, err)

	return errGenerator(FuncName, OperationName, err)
}

// ---------------------------------
