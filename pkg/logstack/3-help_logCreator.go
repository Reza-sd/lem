package logstack

// ---------------------------------
func (f *helperFn[T]) logCreator(errCode T, desArr []any, fnLevel func(string, ...any)) {
	msg, agrs := f.logger.get.msgGenerator(errCode, desArr...)

	if f.logger.get.ifPrintLogsToCli() {
		println()
		fnLevel(msg, agrs)

	}

	if f.logger.get.ifSaveLogsToFile() {
		fnLevel(msg, agrs)
	}
}

// ---------------------------------
