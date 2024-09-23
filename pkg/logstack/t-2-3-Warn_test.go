package logstack

import (
	"fmt"
	"testing"
)

// -----------------------------------------------------
func Test_Warn(t *testing.T) {
	//t.Skip()
	t.Run(`1-RSlogErr`, func(t *testing.T) {
		fnName := "Err.Log()"
		opName := "opName"
		opDes := "opDes"
		errMsg := "error###" //string

		SampleLogger1.Warn.Log(fnName, opName, opDes, errMsg)
		//SampleLogger1.
		//println()
	})
	t.Run(`1-RSlogErr`, func(t *testing.T) {
		fnName := "Err.Log()"
		opName := "opName"
		opDes := "opDes"
		errMsg := fmt.Errorf("miooo") //error type

		SampleLogger1.Warn.Log(fnName, opName, opDes, errMsg)
		//SampleLogger1.
		//println()
	})

	t.Run(`2-RSlogErr`, func(t *testing.T) {
		fnName := "Err.Rlog()"
		opName := "opName"
		opDes := "opDes"
		//errMsg := "error???"
		errCode := errT(10)

		Rerr := SampleLogger1.Warn.Rlog(fnName, opName, opDes, errCode, nil)
		//println("Rerr=",Rerr)
		fmt.Println("Rerr=", Rerr)
		//SampleLogger1.

		//println()
	})

	t.Run(`2-RSlogErr`, func(t *testing.T) {
		fnName := "Err.Rlog()"
		opName := "opName"
		opDes := "opDes"
		//errMsg := "error!!!"
		errCode := errT(12)
		preErrSlice := []errT{10, 11}
		Rerr := SampleLogger1.Warn.Rlog(fnName, opName, opDes, errCode, preErrSlice)
		//println("Rerr=",Rerr)
		fmt.Println("Rerr=", Rerr)
		//SampleLogger1.
		//println()
	})
}

//-----------------------------------------------------
