package logstack

import (
	"fmt"
	"testing"
)

// -----------------------------------------------------
func Test_WarnLogRErrMsg(t *testing.T) {
	//t.Skip()
	t.Run(`1-WarnLogRErrMsg`, func(t *testing.T) {
		fnName := "WarnLogRErrMsg"
		opName := "opName"
		opDes := "opDes"
		err := "error"
		errMsg := SampleLogger.WarnLogRErrMsg(fnName, opName, opDes, err)
		fmt.Printf("errMsg=%v", errMsg)

	})
}

//-----------------------------------------------------
