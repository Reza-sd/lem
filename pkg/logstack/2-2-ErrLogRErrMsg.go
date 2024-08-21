package logstack

// ---------------------------------
func (l *LogCollector) ErrLogRErrMsg(FuncName string, OperationName string, operationDescription string, err any) error {

	l.ErrLog(FuncName, OperationName, operationDescription, err)

	return errGenerator(FuncName, OperationName, err)
}

// ---------------------------------
