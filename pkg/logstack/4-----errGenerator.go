package logstack

import "fmt"

// ----------------------------------------
func errGenerator(FuncName string, Operation string, RetunedError any) error {

	return fmt.Errorf("<==={%v}--%v<---%v[%v]***", pkgName, FuncName, Operation, RetunedError)
	//If error is not nil, it is their responsibility to check for it and handle it (log, return, serve, invoke some retry/cleanup mechanism, etc.).

	//errors.Unwrap
	
}

// ----------------------------------------
