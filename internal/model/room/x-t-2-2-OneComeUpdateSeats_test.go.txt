package room

import (
	"testing"
)

func Test_OneComeUpdateSeats(t *testing.T) {
	//---------------------------------------------
	t.Run("1-return true", func(t *testing.T) {

		Sample_Room_Middle.OneComeUpdateSeats()
		expRoom := *Sample_Room_Middle
		expRoom.UsedSeats = 10

		Assert(t, *Sample_Room_Middle, expRoom)

	})

}
