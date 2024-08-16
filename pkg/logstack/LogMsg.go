package logstack

//import "go/types"

// ----------------------------------------
func LogMsg(packageName string, FuncName string, OperationName string, description string, RetunedError interface{}) (string, string, string, string, string, string, string, string, string) {

	var errStr string
	errStr = "nil"
	//var Rerr error

	switch RetunedError.(type) {
	case string:
		errStr = RetunedError.(string)
	case error:
		if RetunedError.(error) != nil {
			errStr = RetunedError.(string)
		}

	}

	return description, "package", packageName, "func", FuncName, "op", OperationName, "error", errStr

}

//--------------------
