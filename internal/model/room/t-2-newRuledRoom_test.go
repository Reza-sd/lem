package room

import (
	"testing"
)

func Test_newRuledRoom(t *testing.T) {
	//t.Skip()
	t.Run("1-NewRuledRoom", func(t *testing.T) {
		//inp1 := SampleRoom.Middle_Name_3()
		got, _ := NewRuledObjectOFroomT(_START_ROOM_NAME, nil, false)
		got.Act.Print()
		//exp := false
		//logger.Assert(t, got, exp, inp1)
	})
	t.Run("2-NewRuledRoom", func(t *testing.T) {
		got, _ := NewRuledObjectOFroomT(5, nil, true)
		got.Act.Print()

	})
	t.Run("3-NewRuledRoom", func(t *testing.T) {
		got, _ := NewRuledObjectOFroomT(1, nil, false)
		got.Act.Print()

	})
	t.Run("4-NewRuledRoom", func(t *testing.T) {
		got, _ := NewRuledObjectOFroomT(1, []rT{7, 3, 14}, false)
		got.Act.Print()
	})
	t.Run("5-NewRuledRoom exceed maxname", func(t *testing.T) {
		inp1 := rT(_MAX_NAME + 1)
		_, gotErr := NewRuledObjectOFroomT(inp1, []rT{7, 3, 14}, false)
		expErr := []errT{roomTsetT_name_10, NewRuledRoom_10}
		logger.Helper.Assert(t, gotErr, expErr, inp1)
	})

}
