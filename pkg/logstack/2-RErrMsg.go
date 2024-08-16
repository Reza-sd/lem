package logstack

// ---------------------------------
func (l *LogCollector) RErrMsg(FuncName string, OperationName string, RetunedError interface{}) error {
	return errMsg(FuncName, OperationName, RetunedError)
}

// ---------------------------------
