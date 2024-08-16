package logstack

// ---------------------------------
func (l *LogCollector) ErrorToCliFile(FuncName string, OperationName string, err error, operationDescription string) {

	operationDescription = " fail: " + operationDescription

	loggerToCli.Error(LogMsg(l.PackageName, FuncName, OperationName, operationDescription, err))
	loggerToFile.Error(LogMsg(l.PackageName, FuncName, OperationName, operationDescription, err))

}

// ---------------------------------
