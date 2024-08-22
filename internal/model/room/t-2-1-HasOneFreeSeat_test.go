package room

import (
	//"reflect"
	"testing"
)

func Test_HasOneFreeSeat(t *testing.T) {
	//---------------------------------------------
	t.Run("1-return true", func(t *testing.T) {

		got := Sample_Room_Middle.HasOneFreeSeat()
		assert_HasOneFreeSeat(t, got, true)

	})
	//---------------------------------------------
	t.Run("1-return false", func(t *testing.T) {

		Sample_Room_Middle.UsedSeats = Sample_Room_Middle.AllSeats
		got := Sample_Room_Middle.HasOneFreeSeat()
		assert_HasOneFreeSeat(t, got, false)

	})
}

// ==================================
func assert_HasOneFreeSeat(t testing.TB, got, exp bool) {
	t.Helper()
	if got != exp {
		t.Errorf("\n>>>HasOneFreeSeat: \n got=%v<<\n exp=%v<<", got, exp)
	}
}

//======================================
