package logstack

// ---------------------------------
func (e *errLevelT) Log(FuncName string, OperationName string, operationDescription string, errMsg any) {
	//e.logger.
	msg, agrs := LogMsgGenerator(e.logger.get.pkgName(), FuncName, OperationName, operationDescription, errMsg)

	if e.logger.get.ifPrintLogsToCli(){
		println()
		loggerToCli.Error(msg, agrs)
		
	}

	if e.logger.get.ifSaveLogsToFile() {
		loggerToFile.Error(msg, agrs)
	}

}
// ---------------------------------
func (e *errLevelT) Rlog(FuncName string, OperationName string, operationDescription string,errCode errT,previousStatusCodesSlice []errT) []errT {

	e.Log(FuncName, OperationName, operationDescription, e.logger.get.DesForErrCode(errCode))

	//return errGenerator(FuncName, OperationName, err)
	return statwrapper(errCode,previousStatusCodesSlice)
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
