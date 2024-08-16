package logstack

// ----------------WarnLog-----------------
func (l *LogCollector) WarnLog(FuncName string, OperationName string, err interface{}, operationDescription string) {

	if l.LogToCli {
		loggerToCli.Warn(LogMsg(l.PackageName, FuncName, OperationName, operationDescription, err))
	}

	if l.LogToFile {
		loggerToFile.Warn(LogMsg(l.PackageName, FuncName, OperationName, operationDescription, err))

	}

}

//----------------------------------------------
