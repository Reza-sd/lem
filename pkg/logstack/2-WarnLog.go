package logstack

// ----------------WarnLog-----------------
func (l *LogCollector) WarnLog(FuncName string, OperationName string, err interface{}, operationDescription string) {
	s := LogMsg(l.PackageName, FuncName, OperationName, operationDescription, err)
	if l.LogToCli {
		loggerToCli.Warn(s[0], s[1:])
	}

	if l.LogToFile {
		loggerToFile.Warn(s[0], s[1:])

	}

}

//----------------------------------------------
