package logstack

// ---------------------------------
func help_log(l *loggerT, errCode errT, desArr []any, fnLevel func(string, ...any)) {
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
