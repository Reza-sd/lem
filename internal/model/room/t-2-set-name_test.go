package room

import (
	"testing"
)

func Test_Set(t *testing.T) {

	t.Run("1-set", func(t *testing.T) {
		inp1 := RT(2000)
		gotErr := SampleRoom.Start_Name_0().set.name(inp1)
		expErr := []errT{set_name_10}
		logger.Assert(t, gotErr, expErr, inp1)
	})
	t.Run("2-set", func(t *testing.T) {
		inp1 := RT(20)
		gotErr := SampleRoom.Start_Name_0().set.name(inp1)
		var expErr []errT
		logger.Assert(t, gotErr, expErr, inp1)
	})

}
