package logstack

// ----------------WarnLog-----------------
func (l *LogCollector) WarnLog(FuncName string, OperationName string, err interface{}, operationDescription string) {
	msg, agrs := LogMsgGenerator(l.PackageName, FuncName, OperationName, operationDescription, err)
	if l.LogToCli {
		loggerToCli.Warn(msg, agrs)
	}

	if l.LogToFile {
		loggerToFile.Warn(msg, agrs)

	}

}

//----------------------------------------------
