package logstack

// ---------------------------------
func (l *LoggerT) RErrMsg(FuncName string, OperationName string, RetunedError any) error {
	return errGenerator(FuncName, OperationName, RetunedError)
}

// ---------------------------------
