package logstack

// ---------------------------------
func (l *Logger) RErrMsg(FuncName string, OperationName string, RetunedError any) error {
	return errGenerator(FuncName, OperationName, RetunedError)
}

// ---------------------------------
