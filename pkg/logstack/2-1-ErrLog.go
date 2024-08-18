package logstack

// ---------------------------------
func (l *LogCollector) ErrLog(FuncName string, OperationName string,operationDescription string, err interface{} ) {

	msg, agrs := LogMsgGenerator(l.PackageName, FuncName, OperationName, operationDescription, err)
	//msg=fmt.Sprintf("%v-%v", "ErrLog",msg)

	if l.LogToCli {
		loggerToCli.Error(msg, agrs)

	}

	if l.LogToFile {
		loggerToFile.Error(msg, agrs)
	}

}

// ---------------------------------
