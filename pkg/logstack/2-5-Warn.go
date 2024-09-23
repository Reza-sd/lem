package logstack

// ---------------------------------
func (w *warnLevelT) Log(FuncName string, OperationName string, operationDescription string, errMsg any) {
	//e.logger.
	msg, agrs := LogMsgGenerator(w.logger.get.pkgName(), FuncName, OperationName, operationDescription, errMsg)

	if w.logger.get.ifPrintLogsToCli() {
		println()
		loggerToCli.Warn(msg, agrs)

	}

	if w.logger.get.ifSaveLogsToFile() {
		loggerToFile.Warn(msg, agrs)
	}

}

// ---------------------------------
func (w *warnLevelT) Rlog(FuncName string, OperationName string, operationDescription string, errCode errT, previousStatusCodesSlice []errT) []errT {

	w.Log(FuncName, OperationName, operationDescription, w.logger.get.DesForErrCode(errCode))

	return statwrapper(errCode, previousStatusCodesSlice)
}

// ---------------------------------
