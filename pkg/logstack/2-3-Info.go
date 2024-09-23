package logstack

// ---------------------------------
func (i *infoLevelT) Log(FuncName string, OperationName string, operationDescription string, errMsg any) {
	//e.logger.
	msg, agrs := LogMsgGenerator(i.logger.get.pkgName(), FuncName, OperationName, operationDescription, errMsg)

	if i.logger.get.ifPrintLogsToCli() {
		println()
		loggerToCli.Info(msg, agrs)

	}

	if i.logger.get.ifSaveLogsToFile() {
		loggerToFile.Info(msg, agrs)
	}

}

// ---------------------------------
func (i *infoLevelT) Rlog(FuncName string, OperationName string, operationDescription string, errCode errT, previousStatusCodesSlice []errT) []errT {

	i.Log(FuncName, OperationName, operationDescription, i.logger.get.DesForErrCode(errCode))

	return statwrapper(errCode, previousStatusCodesSlice)
}

// ---------------------------------
