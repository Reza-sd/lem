package logstack

import "fmt"

// ----------------------------------------
func errMsg(FuncName string, Operation string, RetunedError interface{}) error {

	//RetunedErrorStr:=fmt.Sprintf("%v",RetunedError)

	//errStr := "<==={" + pkgName + "}--" + FuncName + "<---" + Operation + "[" + RetunedErrorStr + "]***"
	//"<==={%v}--%v<---%v[%v]***"

	return fmt.Errorf("<==={%v}--%v<---%v[%v]***", pkgName, FuncName, Operation, RetunedError)
}

// ----------------------------------------
