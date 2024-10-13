package logstack

// ---------------------------------
func (w *warnLevelT[T]) Log(errCode T, des ...any) {
	help_logCreator(w.logger, errCode, des, loggerToCli.Warn)
}

// ---------------------------------
func (w *warnLevelT[T]) Rlog(errCode T, previousStatusCodesSlice []T, des ...any) []T {

	w.Log(errCode, des...)

	return w.logger.Help.StatWrapper(errCode, previousStatusCodesSlice)
}

// ---------------------------------
