package logstack

import (
	//"fmt"
	"testing"
)

//-----------------------------------------------------
func Test_InfoLog(t *testing.T) {
t.Skip()
	t.Run(`1-InfoLog`, func(t *testing.T) {
		fnName:="InfoLog"
		opName:="opName"
		opDes:="opDes"
		
		 SampleLogger.InfoLog(fnName,opName,opDes)

	})
}
//-----------------------------------------------------