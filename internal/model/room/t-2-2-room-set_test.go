package room

import (
	"testing"
)

func Test_Set(t *testing.T) {

	t.Run("1-set", func(t *testing.T) {
		gotErr := SampleRoom.Start_Name_0().set.name(20000)
		expErr := []errT{set_name_10}
		logger.Assert(t, gotErr, expErr)
	})

}
