package logstack

import (
	"testing"
)

// -----------------------------------------------------
func Test_LogMsgGenerator(t *testing.T) {
	t.Skip()

	t.Run(`1-Print`, func(t *testing.T) {
		SampleLogger.ErrLog("LogMsgGenerator", "OperationName", "errMsg", "operationDescription")
	})
}

//-----------------------------------------------------
