package logstack

// ---------------------------------
func (l *LoggerT) ErrLog(FuncName string, OperationName string, operationDescription string, err any) {

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
/*
logger.Info.Log()
logger.Info.Rlog()
logger.Warn.Log()
logger.Warn.Rlog()
logger.Err.Rlog()
*/
