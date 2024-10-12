package room

import (
	"testing"
)

// ===========================================
func Test_Set(t *testing.T) {

	t.Run("1-set", func(t *testing.T) {
		inp1 := rT(2000)
		gotErr := SampleRoom.Start_Name_0().set.name(inp1)
		expErr := []errT{roomTsetT_name_10}
		logger.Functions.Assert(t, gotErr, expErr, inp1)
	})
	t.Run("2-set", func(t *testing.T) {
		inp1 := rT(20)
		gotErr := SampleRoom.Start_Name_0().set.name(inp1)
		var expErr []errT
		logger.Functions.Assert(t, gotErr, expErr, inp1)
	})

}

//===========================================
