package logstack

// ---------------------------------
func (i *infoLevelT) Log(errCode errT) {

	msg, agrs := i.logger.msgGenerator(errCode)

	if i.logger.get.ifPrintLogsToCli() {
		println()
		loggerToCli.Info(msg, agrs)

	}

	if i.logger.get.ifSaveLogsToFile() {
		loggerToFile.Info(msg, agrs)
	}

}

// ---------------------------------
func (i *infoLevelT) Rlog(errCode errT, previousStatusCodesSlice []errT) []errT {

	i.Log(errCode)

	return i.logger.statWrapper(errCode, previousStatusCodesSlice)
}

// ---------------------------------
