package logstack

// ---------------------------------
func (l *LogCollector) InfoLog(FuncName string, OperationName string, operationDescription string) {

	operationDescription = " success: " + operationDescription

	if l.LogToCli {
		loggerToCli.Info(LogMsg(l.PackageName, FuncName, OperationName, operationDescription, nil))
	}

	if l.LogToFile {
		loggerToFile.Info(LogMsg(l.PackageName, FuncName, OperationName, operationDescription, nil))
	}

}

// ---------------------------------
