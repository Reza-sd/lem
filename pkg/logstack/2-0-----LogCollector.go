package logstack

// ===================struct=========================
//
//	type Attr struct {
//		Key   string
//		Value interface{}
//	}
//
// Debug, Info, Warn, Error, and Panic,
type Logger struct {
	packageName string
	logToFile   bool
	logToCli    bool

	Info info
}
type info struct{
	
}
//=================================================
//var Log LogCollector
