package logstack

import (
	"fmt"
	"testing"
)

// -----------------------------------------------------
func Test_Warn(t *testing.T) {
	//t.Skip()
	t.Run(`1-RSlogErr`, func(t *testing.T) {
		errCode := uint8(10)
		sampleLogger1.Act.Warn.Log(errCode)

	})

	t.Run(`2-RSlogErr`, func(t *testing.T) {
		errCode := uint8(10)

		Rerr := sampleLogger1.Act.Warn.Rlog(errCode, nil)
		fmt.Println("Rerr=", Rerr)

	})

	t.Run(`2-RSlogErr`, func(t *testing.T) {

		errCode := uint8(12)
		preErrSlice := []uint8{10, 11}
		Rerr := sampleLogger1.Act.Warn.Rlog(errCode, preErrSlice)

		fmt.Println("Rerr=", Rerr)

	})
}

//-----------------------------------------------------
