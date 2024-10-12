package room

import (
	"testing"
)

func Test_hasOneFreeSeat(t *testing.T) {
	//t.Skip()
	t.Run("1-hasOneFreeSeat", func(t *testing.T) {
		inp1 := SampleRoom.Middle_Name_3()
		got := inp1.Get.hasOneFreeSeat()
		exp := true
		logger.Function.Assert(t, got, exp, inp1)
	})
	t.Run("2-hasOneFreeSeat", func(t *testing.T) {
		inp1 := SampleRoom.Middle_Name_3()
		got := inp1.Get.hasOneFreeSeat()
		exp := true
		logger.Function.Assert(t, got, exp, inp1)
	})
}
