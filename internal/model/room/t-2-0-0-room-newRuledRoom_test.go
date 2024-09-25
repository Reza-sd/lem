package room

import (
	"testing"
)

func Test_newRuledRoom(t *testing.T) {

	t.Run("1-newPlainRoom", func(t *testing.T) {
		//inp1 := SampleRoom.Middle_Name_3()
		got ,_:= NewRuledRoom(1,1,nil)
		got.Act.Print()
		//exp := false
		//logger.Assert(t, got, exp, inp1)
	})
	// t.Run("1-hasOneFreeSeat", func(t *testing.T) {
	// 	inp1 := SampleRoom.Middle_Name_3()
	// 	got := inp1.Get.hasOneFreeSeat()
	// 	exp := true
	// 	logger.Assert(t, got, exp, inp1)
	// })
}
