package logstack

// ---------------------------------
func (i *infoLevelT[T]) Log(errCode T, des ...any) {

	help_logCreator(i.logger, errCode, des, loggerToCli.Info)

}

// ---------------------------------
func (i *infoLevelT[T]) Rlog(errCode T, previousStatusCodesSlice []T, des ...any) []T {

	i.Log(errCode, des...)

	return i.logger.Help.StatWrapper(errCode, previousStatusCodesSlice)
}

// ---------------------------------
