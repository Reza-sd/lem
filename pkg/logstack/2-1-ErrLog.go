package logstack

// ---------------------------------
func (l *Logger) ErrLog(FuncName string, OperationName string, operationDescription string, err any) {

	msg, agrs := LogMsgGenerator(l.packageName, FuncName, OperationName, operationDescription, err)
	//msg=fmt.Sprintf("%v-%v", "ErrLog",msg)

	if l.logToCli {
		loggerToCli.Error(msg, agrs)

	}

	if l.logToFile {
		loggerToFile.Error(msg, agrs)
	}

}

// ---------------------------------
