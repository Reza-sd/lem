package logstack

import (
	//"fmt"
	"testing"
)

// -----------------------------------------------------
func Test_ErrLog(t *testing.T) {
	//t.Skip()
	t.Run(`1-RSlogErr`, func(t *testing.T) {
		fnName := "Err.Log()"
		opName := "opName"
		opDes := "opDes"
		err := "error###"
		
		SampleLogger1.Err.Log(fnName, opName, opDes, err)
		//SampleLogger1.
	})

	t.Run(`2-RSlogErr`, func(t *testing.T) {
		fnName := "Err.Rlog()"
		opName := "opName"
		opDes := "opDes"
		err := "error***"
		
		SampleLogger1.Err.Rlog(fnName, opName, opDes, err)
		//SampleLogger1.
	})
}

//-----------------------------------------------------
