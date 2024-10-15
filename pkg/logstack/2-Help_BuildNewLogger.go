package logstack

// =================================================

func Help_BuildNewLogger[T errType](packageName string, errCodeDes map[T]string, ifSaveLogsToFile bool, ifPrintLogsToCli bool) *loggerT[T] {

	l := &loggerT[T]{}

	l.data.packageName = packageName
	l.data.ifPrintLogsToCli = ifPrintLogsToCli
	l.data.ifSaveLogsToFile = ifSaveLogsToFile
	l.data.ErrCodeDes = errCodeDes
	l.get.logger = l

	l.Act.Info.logger = l
	l.Act.Warn.logger = l
	l.Act.Err.logger = l

	l.Helper.logger = l

	return l
}

// ===============================================
