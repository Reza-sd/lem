package logstack

// ---------------------------------
func (i *infoLevelT) Log(errCode errT, des ...any) {

	logHelper(i.logger, errCode, des, loggerToCli.Info)
	// msg, agrs := i.logger.get.msgGenerator(errCode, des...)

	// if i.logger.get.ifPrintLogsToCli() {
	// 	println()
	// 	loggerToCli.Info(msg, agrs)

	// }

	// if i.logger.get.ifSaveLogsToFile() {
	// 	loggerToFile.Info(msg, agrs)
	// }

}

// ---------------------------------
func (i *infoLevelT) Rlog(errCode errT, previousStatusCodesSlice []errT, des ...any) []errT {

	i.Log(errCode, des...)

	return i.logger.Help.StatWrapper(errCode, previousStatusCodesSlice)
}

// ---------------------------------
func logHelper(l *loggerT, errCode errT, desArr []any, fnLevel func(string, ...any)) {
	msg, agrs := l.get.msgGenerator(errCode, desArr...)

	if l.get.ifPrintLogsToCli() {
		println()
		fnLevel(msg, agrs)

	}

	if l.get.ifSaveLogsToFile() {
		fnLevel(msg, agrs)
	}
}

// ---------------------------------
