package logstack

import (
	"fmt"
	"testing"
)

// -----------------------------------------------------
func Test_RErrMsg(t *testing.T) {
	t.Skip()
	t.Run(`1-RErrMsg`, func(t *testing.T) {
		fnName := "RErrMsg"
		opName := "opName"

		err := "error"
		errMsg := SampleLogger.RErrMsg(fnName, opName, err)
		fmt.Printf("errMsg=%v", errMsg)

	})
}

//-----------------------------------------------------
