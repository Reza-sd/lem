package logstack

// ---------------------------------
func (l *LogCollector) RErrMsg(FuncName string, OperationName string, RetunedError error) error {
	return errMsg(FuncName, OperationName, RetunedError)
}

// ---------------------------------
