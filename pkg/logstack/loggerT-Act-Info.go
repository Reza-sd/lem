package logstack

// ---------------------------------
func (i *infoLevelT[T]) Log(errCode T, des ...any) {

	i.logger.Helper.logCreator(errCode, des, loggerToCli.Error)

}

// ---------------------------------
func (i *infoLevelT[T]) Rlog(errCode T, previousStatusCodesSlice []T, des ...any) []T {

	i.Log(errCode, des...)

	return i.logger.Helper.StatWrapper(errCode, previousStatusCodesSlice)
}

// ---------------------------------
