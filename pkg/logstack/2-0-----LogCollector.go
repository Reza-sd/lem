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
*/
type Logger struct {
	data data
	get  getter

	Info info
	Warn warn
	Fail fail //

}
type data struct {
	packageName string
	logsToFile  bool
	logsToCli   bool
}
type info struct {
	logger *Logger
}
type warn struct {
	logger *Logger
}
type fail struct {
	logger *Logger
}
type getter struct {
	logger *Logger
}

//=================================================
func (get *getter) pkgName() string {
	return get.logger.data.packageName
}
func (get *getter) ifSaveLogsToFile() bool {
	return get.logger.data.logsToFile
}
func (get *getter) ifPrintLogsToCli() bool {
	return get.logger.data.logsToCli
}

//=================================================
//var Log LogCollector
