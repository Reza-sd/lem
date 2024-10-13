package logstack

// ---------------------------------
func (e *errLevelT) Log(errCode errT, des ...any) {

	help_logCreator(e.logger, errCode, des, loggerToCli.Error)

}

// ---------------------------------
func (e *errLevelT) Rlog(errCode errT, previousStatusCodesSlice []errT, des ...any) []errT {

	e.Log(errCode, des...)

	return e.logger.Help.StatWrapper(errCode, previousStatusCodesSlice)
}

// ---------------------------------
