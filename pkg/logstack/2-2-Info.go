package logstack

// ---------------------------------
func (i *infoLevelT) Log(errCode errT, des ...any) {

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
func (i *infoLevelT) Rlog(errCode errT, previousStatusCodesSlice []errT, des ...any) []errT {

	i.Log(errCode, des...)

	return i.logger.StatWrapper(errCode, previousStatusCodesSlice)
}

// ---------------------------------
