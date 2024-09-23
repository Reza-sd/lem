package logstack

import (
	//"fmt"
	"testing"
)

// -----------------------------------------------------
func Test_Warn(t *testing.T) {
	//t.Skip()
	t.Run(`1-Warn.log`, func(t *testing.T) {
		// fnName := "func1"
		// opName := "opName"
		// opDes := "opDes"
		// errMsg := "error###" //string
		errCode := errT(10)

		SampleLogger1.Warn.Log(errCode)

	})
	// t.Run(`1-Warn.log`, func(t *testing.T) {
	// 	fnName := "func1"
	// 	opName := "opName"
	// 	opDes := "opDes"
	// 	errMsg := fmt.Errorf("miooo") //error type

	// 	SampleLogger1.Warn.Log(fnName, opName, opDes, errMsg)

	// })

	// t.Run(`2-Warn.Rlog`, func(t *testing.T) {
	// 	fnName := "func1"
	// 	opName := "opName"
	// 	opDes := "opDes"
	// 	errCode := errT(10)

	// 	Rerr := SampleLogger1.Warn.Rlog(fnName, opName, opDes, errCode, nil)
	// 	fmt.Println("Rerr=", Rerr)

	// })

	// t.Run(`2-Warn.Rlog`, func(t *testing.T) {
	// 	fnName := "func1"
	// 	opName := "opName"
	// 	opDes := "opDes"
	// 	errCode := errT(12)
	// 	preErrSlice := []errT{10, 11}
	// 	Rerr := SampleLogger1.Warn.Rlog(fnName, opName, opDes, errCode, preErrSlice)
	// 	fmt.Println("Rerr=", Rerr)

	// })
}

//-----------------------------------------------------
