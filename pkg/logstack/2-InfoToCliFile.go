package logstack

// ---------------------------------
func (l *LogCollector) InfoToCliFile(FuncName string, OperationName string, operationDescription string) {

	operationDescription = " success: " + operationDescription

	loggerToCli.Info(LogMsg(l.PackageName, FuncName, OperationName, operationDescription, nil))

	loggerToFile.Info(LogMsg(l.PackageName, FuncName, OperationName, operationDescription, nil))
}

// ---------------------------------
