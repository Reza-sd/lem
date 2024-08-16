package logstack

import "fmt"

//import "go/types"

// ----------------------------------------
func LogMsg(packageName string, FuncName string, OperationName string, description string, RetunedError interface{}) (string, string, string, string, string, string, string, string, string) {

	errString := fmt.Sprintf("%v", RetunedError)

	return description, "package", packageName, "func", FuncName, "op", OperationName, "error", errString

}

//----------------------------------------------
