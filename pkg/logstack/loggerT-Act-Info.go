package logstack

// ---------------------------------
func (i *infoLevelT) Log(errCode errT, des ...any) {

	help_log(i.logger, errCode, des, loggerToCli.Info)

}

// ---------------------------------
func (i *infoLevelT) Rlog(errCode errT, previousStatusCodesSlice []errT, des ...any) []errT {

	i.Log(errCode, des...)

	return i.logger.Help.StatWrapper(errCode, previousStatusCodesSlice)
}

// ---------------------------------
