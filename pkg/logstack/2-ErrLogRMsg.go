package logstack

// ---------------------------------
func (l *LogCollector) ErrLogRMsg(FuncName string, OperationName string, err interface{}, operationDescription string) error {

	l.ErrLog(FuncName, OperationName, err, operationDescription)

	return errGenerator(FuncName, OperationName, err) //optional return
}

// ---------------------------------
