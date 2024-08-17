package logstack

// ---------------------------------
func (l *LogCollector) InfoLog(FuncName string, OperationName string, operationDescription string) {
	msg, agrs := LogMsgGenerator(l.PackageName, FuncName, OperationName, operationDescription, nil)
	if l.LogToCli {
		loggerToCli.Info(msg, agrs)
	}

	if l.LogToFile {
		loggerToFile.Info(msg, agrs)
	}

}

// ---------------------------------
