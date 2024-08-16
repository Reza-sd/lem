package logstack

// ----------------WarnToCliFile-----------------
func (l *LogCollector) WarnToCliFile(FuncName string, OperationName string, err error, operationDescription string) {

	operationDescription = " !!!: " + operationDescription

	loggerToCli.Warn(LogMsg(l.PackageName, FuncName, OperationName, operationDescription, err))
	loggerToFile.Warn(LogMsg(l.PackageName, FuncName, OperationName, operationDescription, err))

}

//----------------------------------------------
