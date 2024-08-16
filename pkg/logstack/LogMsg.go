package logstack

import "fmt"

//import "go/types"

// ----------------------------------------
func LogMsg(packageName string, FuncName string, OperationName string, description string, RetunedError interface{}) (string, string, string, string, string, string, string, string, string) {

	
	//errString := "nil"
	errString:=fmt.Sprintf("%v",RetunedError)

	// switch v := RetunedError.(type) {
	// case string:
	// 	if v!=""{
	// 		errString = v
	// 	}
		
	// case error:
	// 	if v != nil {
	// 		errString = v.Error()
	// 	} 

	// }

	return description, "package", packageName, "func", FuncName, "op", OperationName, "error", errString

}

//----------------------------------------------
