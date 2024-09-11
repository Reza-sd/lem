package logstack

// ===================struct=========================
// type Attr struct {
// 	Key   string
// 	Value interface{}
// }

type LogCollector struct {
	PackageName string
	LogToFile   bool
	LogToCli    bool
}

//=================================================
var Log LogCollector