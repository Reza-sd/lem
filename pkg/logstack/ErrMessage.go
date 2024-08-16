package logstack

import "fmt"

// ----------------------------------------
func errMsg(FuncName string, Operation string, RetunedError interface{}) error {

	return fmt.Errorf("<==={%v}--%v<---%v[%v]***", pkgName, FuncName, Operation, RetunedError)
}

// ----------------------------------------
