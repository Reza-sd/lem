package logstack

import (
	"fmt"
	"testing"
)

// -----------------------------------------------------
func Test_Err(t *testing.T) {
	//t.Skip()
	t.Run(`1-RSlogErr`, func(t *testing.T) {
		errCode := errT(10)
		SampleLogger1.Err.Log(errCode)

	})

	t.Run(`2-RSlogErr`, func(t *testing.T) {
		errCode := errT(10)

		Rerr := SampleLogger1.Err.Rlog(errCode, nil)
		fmt.Println("Rerr=", Rerr)

	})

	t.Run(`2-RSlogErr`, func(t *testing.T) {

		errCode := errT(12)
		preErrSlice := []errT{10, 11}
		Rerr := SampleLogger1.Err.Rlog(errCode, preErrSlice)

		fmt.Println("Rerr=", Rerr)

	})
}

//-----------------------------------------------------
