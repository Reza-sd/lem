package logstack

// =================================================

func BuildNewLogger(packageName string, errCodeDes map[errT]string, ifSaveLogsToFile bool, ifPrintLogsToCli bool) *loggerT {
	l := &loggerT{}
	// l.data = struct{
	// 	packageName:      packageName,
	// 	ifSaveLogsToFile: ifSaveLogsToFile,
	// 	ifPrintLogsToCli: ifPrintLogsToCli,
	// 	ErrCodeDes:       errCodeDes,
	// }
	l.data.packageName = packageName
	l.data.ifPrintLogsToCli = ifPrintLogsToCli
	l.data.ifSaveLogsToFile = ifSaveLogsToFile
	l.data.ErrCodeDes = errCodeDes
	l.get.logger = l
	//l.Act.logger=l

	// l.Info.logger = l
	l.Act.Info.logger = l
	l.Act.Warn.logger = l
	l.Act.Err.logger = l

	// l.Warn.logger = l
	// l.Err.logger = l

	return l
}

// ===============================================
