package logstack

// ---------------------------------
func (w *warnLevelT) Log(errCode errT) {
	//e.logger.
	msg, agrs := w.logger.msgGenerator(errCode)

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

	w.Log(errCode)

	return statWrapper(errCode, previousStatusCodesSlice)
}

// ---------------------------------
