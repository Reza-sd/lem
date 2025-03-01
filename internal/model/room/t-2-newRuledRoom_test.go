package room

import (
	"testing"
)

func Test_newRuledRoom(t *testing.T) {
	//t.Skip()
	t.Run("1-NewRuledRoom", func(t *testing.T) {
		//inp1 := SampleRoom.Middle_Name_3()
		got, _ := Return_NewRuledObjectOFroomT(_START_ROOM_NAME, nil, false)
		got.Act.Print()
		//exp := false
		//logger.Assert(t, got, exp, inp1)
	})
	t.Run("2-NewRuledRoom", func(t *testing.T) {
		got, _ := Return_NewRuledObjectOFroomT(5, nil, true)
		got.Act.Print()

	})
	t.Run("3-NewRuledRoom", func(t *testing.T) {
		got, _ := Return_NewRuledObjectOFroomT(1, nil, false)
		got.Act.Print()

	})
	t.Run("4-NewRuledRoom", func(t *testing.T) {
		got, _ := Return_NewRuledObjectOFroomT(1, []rT{7, 3, 14}, false)
		got.Act.Print()
	})
	t.Run("5-NewRuledRoom exceed maxname", func(t *testing.T) {
		inp1 := rT(_MAX_NAME + 1)
		_, gotErr := Return_NewRuledObjectOFroomT(inp1, []rT{7, 3, 14}, false)
		expErr := []errT{_ERR_roomT_setT_name_10, _ERR_NewRuledRoom_10}
		logger.Helper.Assert(t, gotErr, expErr, inp1)
	})

}
