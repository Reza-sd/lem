package logstack

import (
	//"fmt"
	"testing"
)

// -----------------------------------------------------
func Test_ErrLog(t *testing.T) {
	t.Skip()
	t.Run(`1-RSlogErr`, func(t *testing.T) {
		fnName := "ErrLog"
		opName := "opName"
		opDes := "opDes"
		err := "error"
		SampleLogger.ErrLog(fnName, opName, opDes, err)

	})
}

//-----------------------------------------------------
