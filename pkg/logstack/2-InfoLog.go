package logstack

// ---------------------------------
func (l *LogCollector) InfoLog(FuncName string, OperationName string, operationDescription string) {
	s := LogMsg(l.PackageName, FuncName, OperationName, operationDescription, nil)
	if l.LogToCli {
		loggerToCli.Info(s[0], s[1:])
	}

	if l.LogToFile {
		loggerToFile.Info(s[0], s[1:])
	}

}

// ---------------------------------
