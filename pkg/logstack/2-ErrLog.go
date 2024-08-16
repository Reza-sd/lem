package logstack

// ---------------------------------
func (l *LogCollector) ErrLog(FuncName string, OperationName string, err interface{}, operationDescription string) {
	s := LogMsg(l.PackageName, FuncName, OperationName, operationDescription, err)

	if l.LogToCli {
		loggerToCli.Error(s[0], s[1:])
	}

	if l.LogToFile {
		loggerToFile.Error(s[0], s[1:])
	}

}

// ---------------------------------
