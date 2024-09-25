package room

import (
	"testing"
)

func Test_newRuledRoom(t *testing.T) {

	t.Run("1-NewRuledRoom", func(t *testing.T) {
		//inp1 := SampleRoom.Middle_Name_3()
		got, _ := NewRuledRoom(startRoomName, 10, nil)
		got.Act.Print()
		//exp := false
		//logger.Assert(t, got, exp, inp1)
	})
	t.Run("2-NewRuledRoom", func(t *testing.T) {
		got, _ := NewRuledRoom(startRoomName, 0, nil)
		got.Act.Print()

	})
	t.Run("3-NewRuledRoom", func(t *testing.T) {
		got, _ := NewRuledRoom(1, 0, nil)
		got.Act.Print()

	})
	t.Run("4-NewRuledRoom", func(t *testing.T) {
		got, _ := NewRuledRoom(1, 0, []RT{7, 3, 14})
		got.Act.Print()
	})

}
