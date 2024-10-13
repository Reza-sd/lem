package logstack

// ---------------------------------
func (i *infoLevelT[T]) Log(errCode T, des ...any) {

	i.logger.Help.logCreator(errCode, des, loggerToCli.Error)

}

// ---------------------------------
func (i *infoLevelT[T]) Rlog(errCode T, previousStatusCodesSlice []T, des ...any) []T {

	i.Log(errCode, des...)

	return i.logger.Help.StatWrapper(errCode, previousStatusCodesSlice)
}

// ---------------------------------
