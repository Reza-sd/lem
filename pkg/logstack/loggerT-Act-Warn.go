package logstack

// ---------------------------------
func (w *warnLevelT) Log(errCode errT, des ...any) {
	//e.logger.
	msg, agrs := w.logger.get.msgGenerator(errCode, des...)

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

	return w.logger.Function.StatWrapper(errCode, previousStatusCodesSlice)
}

// ---------------------------------
