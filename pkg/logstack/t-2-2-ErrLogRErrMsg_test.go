package logstack

import (
	"fmt"
	"testing"
)

// -----------------------------------------------------
func Test_ErrLogRErrMsg(t *testing.T) {
	t.Skip()
	t.Run(`1-ErrLogRErrMsg`, func(t *testing.T) {
		fnName := "ErrLogRErrMsg"
		opName := "opName"
		opDes := "opDes"
		err := "error"
		errMsg := SampleLogger.ErrLogRErrMsg(fnName, opName, opDes, err)
		fmt.Printf("errMsg=%v", errMsg)

	})
}

//-----------------------------------------------------
