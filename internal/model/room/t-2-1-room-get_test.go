package room

import (
	//"fmt"
	"testing"
)

func Test_hasOneFreeSeat(t *testing.T) {

	t.Run("1-hasOneFreeSeat", func(t *testing.T) {
		inp1 := SampleRoom.Middle_Name_3()
		got := inp1.Get.hasOneFreeSeat()
		exp := false
		//fmt.Println(inp)
		logger.Assert(t, got, exp, inp1)
	})
	t.Run("1-hasOneFreeSeat", func(t *testing.T) {
		inp1 := SampleRoom.Middle_Name_3()

		got := inp1.Get.hasOneFreeSeat()
		exp := true
		//fmt.Println(inp)
		logger.Assert(t, got, exp, inp1)
	})
}
