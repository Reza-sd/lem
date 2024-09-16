package logstack

// ----------------WarnLog-----------------
func (l *LoggerT) WarnLog(FuncName string, OperationName string, operationDescription string, err any) {
	msg, agrs := LogMsgGenerator(l.packageName, FuncName, OperationName, operationDescription, err)
	if l.logToCli {
		loggerToCli.Warn(msg, agrs)
	}

	if l.logToFile {
		loggerToFile.Warn(msg, agrs)

	}

}

//----------------------------------------------
