package room

import (
	"testing"
)

func Test_newRuledRoom(t *testing.T) {
	//t.Skip()
	t.Run("1-NewRuledRoom", func(t *testing.T) {
		//inp1 := SampleRoom.Middle_Name_3()
		got, _ := NewRuledRoom(startRoomName, nil, false)
		got.Act.Print()
		//exp := false
		//logger.Assert(t, got, exp, inp1)
	})
	t.Run("2-NewRuledRoom", func(t *testing.T) {
		got, _ := NewRuledRoom(5, nil, true)
		got.Act.Print()

	})
	t.Run("3-NewRuledRoom", func(t *testing.T) {
		got, _ := NewRuledRoom(1, nil, false)
		got.Act.Print()

	})
	t.Run("4-NewRuledRoom", func(t *testing.T) {
		got, _ := NewRuledRoom(1, []RT{7, 3, 14}, false)
		got.Act.Print()
	})
	t.Run("5-NewRuledRoom", func(t *testing.T) {
		inp1 := RT(maxName + 1)
		_, gotErr := NewRuledRoom(inp1, []RT{7, 3, 14}, false)
		expErr := []errT{set_name_10, NewRuledRoom_10}
		logger.Assert(t, gotErr, expErr, inp1)
	})

}
