package logstack

// ---------------------------------
func (l *LogCollector) ErrLogRMsg(FuncName string, OperationName string, err interface{}, operationDescription string) error {

	l.ErrLog(FuncName, OperationName, operationDescription,err)

	return errGenerator(FuncName, OperationName, err) //optional return
}

// ---------------------------------
