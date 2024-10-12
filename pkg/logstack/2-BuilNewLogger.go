package logstack



// =================================================

func BuildNewLogger(packageName string, errCodeDes map[errT]string, ifSaveLogsToFile bool, ifPrintLogsToCli bool) *loggerT {
	l := &loggerT{}
	l.data = data{
		packageName:      packageName,
		ifSaveLogsToFile: ifSaveLogsToFile,
		ifPrintLogsToCli: ifPrintLogsToCli,
		ErrCodeDes:       errCodeDes,
	}
	l.get.logger = l
	l.Info.logger = l
	l.Warn.logger = l
	l.Err.logger = l

	return l
}

// ===============================================