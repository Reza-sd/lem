package logstack

import "fmt"

// ----------------------------------------
func errGenerator(FuncName string, Operation string, RetunedError any) error {
	RetunedErrorErrorType := fmt.Errorf("%v", RetunedError) //?????

	// wrapping RetunedError using %w format specifier and adding more context

	return fmt.Errorf("<==={%v}--%v<---%v[%w]***", pkgName, FuncName, Operation, RetunedErrorErrorType)
	//If error is not nil, it is their responsibility to check for it and handle it (log, return, serve, invoke some retry/cleanup mechanism, etc.).

	//errors.Unwrap
}

// ----------------------------------------
