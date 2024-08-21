package logstack

// ---------------------------------
func (l *LogCollector) RErrMsg(FuncName string, OperationName string, RetunedError any) error {
	return errGenerator(FuncName, OperationName, RetunedError)
}

// ---------------------------------
