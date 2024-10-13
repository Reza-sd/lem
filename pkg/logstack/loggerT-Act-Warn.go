package logstack

// ---------------------------------
func (w *warnLevelT) Log(errCode errT, des ...any) {
	help_log(w.logger, errCode, des, loggerToCli.Warn)
}

// ---------------------------------
func (w *warnLevelT) Rlog(errCode errT, previousStatusCodesSlice []errT, des ...any) []errT {

	w.Log(errCode, des...)

	return w.logger.Help.StatWrapper(errCode, previousStatusCodesSlice)
}

// ---------------------------------
