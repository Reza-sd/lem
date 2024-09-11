package logstack

// ---------------------------------
func (l *Logger) ErrLogRErrMsg(FuncName string, OperationName string, operationDescription string, err any) error {

	l.ErrLog(FuncName, OperationName, operationDescription, err)

	return errGenerator(FuncName, OperationName, err)
}

// ---------------------------------
