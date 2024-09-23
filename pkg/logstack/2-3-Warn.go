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
func (w *warnLevelT) Rlog(errCode errT, previousStatusCodesSlice []errT) []errT {

	w.Log(errCode)

	return w.logger.StatWrapper(errCode, previousStatusCodesSlice)
}

// ---------------------------------
