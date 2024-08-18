package logstack

import (
	"fmt"
	"testing"
)

//-----------------------------------------------------
func Test_errGenerator(t *testing.T) {
	t.Skip()

	t.Run(`1-Print`, func(t *testing.T) {
		errMsg:= errGenerator("errGenerator","returnErr","prevouis error")
		fmt.Printf("%v\n",errMsg)
		//SampleLogger.ErrLog("LogMsgGenerator", "OperationName", "errMsg", "operationDescription")
	})
}
//-----------------------------------------------------