package logstack

// ---------------------------------
func (e *errLevelT) Log(FuncName string, OperationName string, operationDescription string, err any) {
	//e.logger.
	msg, agrs := LogMsgGenerator(e.logger.data.packageName, FuncName, OperationName, operationDescription, err)

	if e.logger.get.ifPrintLogsToCli(){
		loggerToCli.Error(msg, agrs)

	}

	if e.logger.get.ifSaveLogsToFile() {
		loggerToFile.Error(msg, agrs)
	}

}
// ---------------------------------
func (e *errLevelT) Rlog(FuncName string, OperationName string, operationDescription string, err any) error {

	e.Log(FuncName, OperationName, operationDescription, err)

	return errGenerator(FuncName, OperationName, err)
}

// ---------------------------------
// ---------------------------------
/*
logger.Info.Log()
logger.Info.Rlog()
logger.Warn.Log()
logger.Warn.Rlog()
logger.Err.Rlog()
*/
