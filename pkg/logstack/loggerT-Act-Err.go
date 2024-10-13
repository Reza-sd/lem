package logstack

// ---------------------------------
func (e *errLevelT[T]) Log(errCode T, des ...any) {

	e.logger.Helper.logCreator(errCode, des, loggerToCli.Error)

}

// ---------------------------------
func (e *errLevelT[T]) Rlog(errCode T, previousStatusCodesSlice []T, des ...any) []T {

	e.Log(errCode, des...)

	return e.logger.Helper.StatWrapper(errCode, previousStatusCodesSlice)
}

// ---------------------------------
