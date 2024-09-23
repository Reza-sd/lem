package logstack

// ---------------------------------
func (e *errLevelT) Log(errCode errT) {

	msg, agrs := e.logger.msgGenerator(errCode)

	if e.logger.get.ifPrintLogsToCli() {
		println()
		loggerToCli.Error(msg, agrs)

	}

	if e.logger.get.ifSaveLogsToFile() {
		loggerToFile.Error(msg, agrs)
	}

}

// ---------------------------------
func (e *errLevelT) Rlog(errCode errT, previousStatusCodesSlice []errT) []errT {

	e.Log(errCode)

	return e.logger.StatWrapper(errCode, previousStatusCodesSlice)
}

// ---------------------------------
