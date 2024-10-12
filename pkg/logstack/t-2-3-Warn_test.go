package logstack

import (
	"fmt"
	"testing"
)

// -----------------------------------------------------
func Test_Warn(t *testing.T) {
	//t.Skip()
	t.Run(`1-RSlogErr`, func(t *testing.T) {
		errCode := errT(10)
		sampleLogger1.Act.Warn.Log(errCode)

	})

	t.Run(`2-RSlogErr`, func(t *testing.T) {
		errCode := errT(10)

		Rerr := sampleLogger1.Act.Warn.Rlog(errCode, nil)
		fmt.Println("Rerr=", Rerr)

	})

	t.Run(`2-RSlogErr`, func(t *testing.T) {

		errCode := errT(12)
		preErrSlice := []errT{10, 11}
		Rerr := sampleLogger1.Act.Warn.Rlog(errCode, preErrSlice)

		fmt.Println("Rerr=", Rerr)

	})
}

//-----------------------------------------------------
