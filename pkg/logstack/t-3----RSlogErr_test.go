package logstack

import (
	"fmt"
	"testing"
)

//-----------------------------------------------------
func Test_RSlogErr(t *testing.T) {
t.Skip()
	t.Run(`1-RSlogErr`, func(t *testing.T) {
		errMsg:= RSlogErr("logstackpk","RSlogErr-fn","RSlogErr","return slog when logger stack fail","error")
		fmt.Printf("%v\n",errMsg)
		//SampleLogger.ErrLog("LogMsgGenerator", "OperationName", "errMsg", "operationDescription")
	})
}
//-----------------------------------------------------