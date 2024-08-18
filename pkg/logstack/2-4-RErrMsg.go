package logstack

// ---------------------------------
func (l *LogCollector) RErrMsg(FuncName string, OperationName string, RetunedError interface{}) error {
	return errGenerator(FuncName, OperationName, RetunedError)
}

// ---------------------------------
