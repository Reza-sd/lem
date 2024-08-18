package logstack

import (
	//"fmt"
	"testing"
)

// -----------------------------------------------------
func Test_WarnLog(t *testing.T) {
	t.Skip()
	t.Run(`1-WarnLog`, func(t *testing.T) {
		fnName := "WarnLog"
		opName := "opName"
		opDes := "opDes"
		err := "error"
		SampleLogger.WarnLog(fnName, opName, opDes, err)

	})
}

//-----------------------------------------------------
