package logstack

// ---------------------------------
func (l *LogCollector) ErrLog(FuncName string, OperationName string, err interface{}, operationDescription string) {
	msg, agrs := LogMsgGenerator(l.PackageName, FuncName, OperationName, operationDescription, err)

	if l.LogToCli {
		loggerToCli.Error(msg, agrs)
		
	}

	if l.LogToFile {
		loggerToFile.Error(msg, agrs)
	}

}

// ---------------------------------
