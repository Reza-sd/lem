package logstack

// ---------------------------------
func (e *errLevelT) Log(errCode errT, des ...any) {

	msg, agrs := e.logger.msgGenerator(errCode, des...)

	if e.logger.get.ifPrintLogsToCli() {
		println()
		loggerToCli.Error(msg, agrs)

	}

	if e.logger.get.ifSaveLogsToFile() {
		loggerToFile.Error(msg, agrs)
	}

}

// ---------------------------------
func (e *errLevelT) Rlog(errCode errT, previousStatusCodesSlice []errT, des ...any) []errT {

	e.Log(errCode, des...)

	return e.logger.StatWrapper(errCode, previousStatusCodesSlice)
}

// ---------------------------------
