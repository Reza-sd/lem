package logstack

import "fmt"

//import "go/types"

// ----------------------------------------
func LogMsg(packageName string, FuncName string, OperationName string, description string, RetunedError interface{}) []string {

	//errString := fmt.Sprintf("%v", RetunedError)

	return []string{description, "package", packageName, "func", FuncName, "op", OperationName, "error", fmt.Sprintf("%v", RetunedError)}

}

//----------------------------------------------
