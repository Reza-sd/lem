package logstack

import (
	//"fmt"
	"testing"
)

// -----------------------------------------------------
func Test_Info(t *testing.T) {
	//t.Skip()
	t.Run(`1-Info`, func(t *testing.T) {
		// fnName := "func1"
		// opName := "opName"
		// opDes := "opDes"
		// errMsg := "error###" //string
		errCode := errT(10)

		SampleLogger1.Info.Log(errCode)

	})
	// t.Run(`1-Info`, func(t *testing.T) {
	// 	fnName := "func1"
	// 	opName := "opName"
	// 	opDes := "opDes"
	// 	errMsg := fmt.Errorf("miooo") //error type

	// 	SampleLogger1.Info.Log(fnName, opName, opDes, errMsg)

	// })

	// t.Run(`2-Info`, func(t *testing.T) {
	// 	fnName := "func1"
	// 	opName := "opName"
	// 	opDes := "opDes"
	// 	errCode := errT(10)

	// 	Rerr := SampleLogger1.Info.Rlog(fnName, opName, opDes, errCode, nil)
	// 	fmt.Println("Rerr=", Rerr)

	// })

	// t.Run(`2-Info`, func(t *testing.T) {
	// 	fnName := "func1"
	// 	opName := "opName"
	// 	opDes := "opDes"
	// 	errCode := errT(12)
	// 	preErrSlice := []errT{10, 11}
	// 	Rerr := SampleLogger1.Info.Rlog(fnName, opName, opDes, errCode, preErrSlice)
	// 	fmt.Println("Rerr=", Rerr)

	// })
}

//-----------------------------------------------------
