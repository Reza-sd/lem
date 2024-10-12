package logstack

// ===================struct=========================
//
//	type Attr struct {
//		Key   string
//		Value interface{}
//	}
//
// Debug, Info, Warn, Error, and Panic,
/*
Log Levels in slog
1-Debug: For debug-level messages, typically used for detailed information useful for debugging.
2-Info: For informational messages, usually to signify the normal operation of the application.
3-Warn: For warning messages, which indicate a potential problem or something to pay attention to.
4-Error: For error messages, representing an issue that caused the operation to fail.
5-Fatal: For fatal messages, which indicate a critical error after which the program cannot continue. Typically, the application is terminated after logging a fatal message.
//-------------------------------
Debug: Lowest log level, used for detailed debugging information.

Info: General information about the application's behavior.

Warn: Indicates potential issues or unexpected events.

Error: Serious errors that may prevent the application from functioning correctly.

Panic: Critical errors that should cause the application to crash.
//--------------------------
status codes: are used to communicate the result of an operation.
error= invalid condition
*/
type loggerT struct {
	data struct {
		packageName      string
		ifSaveLogsToFile bool
		ifPrintLogsToCli bool
		ErrCodeDes       map[errT]string
	}

	get getter

	Act struct {
		//logger *loggerT

		Info infoLevelT //e.g: normal status code
		Warn warnLevelT
		Err  errLevelT //unexpected status code
	}

	Functions function
	// error means violation in business logic
	// in software development, an "error" often refers to a violation of business logic.
	// unexpected = Must not happen
	//Stat stat //http 400
	//is it invalid but expected
	// invalid and unexpected
	// valid and expedted
	// valid and unexpected
	//error == 100% invalid and/or 100% unexpected and/or 100% violation of func operation.

	// workflow
}

/*

    logger.Debug("This is a debug message")
    logger.Info("This is an info message")
    logger.Warn("This is a warning message")
    logger.Error("This is an error message")
    logger.Crit("This is a critical error message")

logger.Info.Log()
logger.Info.Rlog()
logger.Warn.Log()
logger.Warn.Rlog()
logger.Err.log()
logger.Err.Rlog() //(Rlog:return error + log)
*/

//	type data struct {
//		packageName      string
//		ifSaveLogsToFile bool
//		ifPrintLogsToCli bool
//		ErrCodeDes       map[errT]string
//	}
type infoLevelT struct {
	logger *loggerT
}
type warnLevelT struct {
	logger *loggerT
}
type errLevelT struct {
	logger *loggerT
}

type getter struct {
	logger *loggerT
}

// type action struct{
// 	//logger *loggerT

// 	Info infoLevelT //e.g: normal status code
// 	Warn warnLevelT
// 	Err  errLevelT //unexpected status code
// }

type function struct {
}

// =================================================
func (get *getter) pkgName() string {
	return get.logger.data.packageName
}
func (get *getter) ifSaveLogsToFile() bool {
	return get.logger.data.ifSaveLogsToFile
}
func (get *getter) ifPrintLogsToCli() bool {
	return get.logger.data.ifPrintLogsToCli
}
func (get *getter) desForErrCode(CodeNumber errT) string {

	st, ok := get.logger.data.ErrCodeDes[CodeNumber]
	if ok {
		return st
	}
	return ""

}
