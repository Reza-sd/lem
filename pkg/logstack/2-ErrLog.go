package logstack

// ---------------------------------
func (l *LogCollector) ErrLog(FuncName string, OperationName string, err interface{}, operationDescription string) {

	if l.LogToCli {
		loggerToCli.Error(LogMsg(l.PackageName, FuncName, OperationName, operationDescription, err))
	}

	if l.LogToFile {
		loggerToFile.Error(LogMsg(l.PackageName, FuncName, OperationName, operationDescription, err))
	}

}

// ---------------------------------
