package logstack

// ---------------------------------
func (i *infoLevelT) Log(FuncName string, OperationName string, operationDescription string) {
	logger = i.logger
	msg, agrs := LogMsgGenerator(logger.packageName, FuncName, OperationName, operationDescription, nil)
	if logger.logToCli {
		loggerToCli.Info(msg, agrs)
	}

	if logger.logToFile {
		loggerToFile.Info(msg, agrs)
	}

}

// ---------------------------------
