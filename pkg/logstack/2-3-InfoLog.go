package logstack

// ---------------------------------
func (l *Logger) InfoLog(FuncName string, OperationName string, operationDescription string) {
	msg, agrs := LogMsgGenerator(l.packageName, FuncName, OperationName, operationDescription, nil)
	if l.logToCli {
		loggerToCli.Info(msg, agrs)
	}

	if l.logToFile {
		loggerToFile.Info(msg, agrs)
	}

}

// ---------------------------------
