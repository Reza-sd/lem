package logstack

// ---------------------------------
func (w *warnLevelT) Log(errCode errT, des ...any) {
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
func (w *warnLevelT) Rlog(errCode errT, previousStatusCodesSlice []errT, des ...any) []errT {

	w.Log(errCode, des...)

	return w.logger.StatWrapper(errCode, previousStatusCodesSlice)
}

// ---------------------------------
