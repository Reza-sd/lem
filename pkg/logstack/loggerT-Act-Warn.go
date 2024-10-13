package logstack

// ---------------------------------
func (w *warnLevelT[T]) Log(errCode T, des ...any) {

	w.logger.Helper.logCreator(errCode, des, loggerToCli.Error)

}

// ---------------------------------
func (w *warnLevelT[T]) Rlog(errCode T, previousStatusCodesSlice []T, des ...any) []T {

	w.Log(errCode, des...)

	return w.logger.Helper.StatWrapper(errCode, previousStatusCodesSlice)
}

// ---------------------------------
