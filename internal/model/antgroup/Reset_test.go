package antgroup

import (
	//"reflect"
	//"fmt"
	"testing"
)

func Test_ResetFartory(t *testing.T) {
	t.Skip()
	//------------
	t.Run(`1-return an error when numberOfAnts=0`, func(t *testing.T) {
		//---Arrange---

		myAntGroup := mySampleAntGroup.Ant_2_initmode_room_0()
		myAntGroup.CurrentSequenceNumber = 10
		//myAntGroup.Print()
		//---Act---
		myAntGroup.Reset()
		//---Assert----
		assert_If_Two_AntGroup_Same(t, myAntGroup, mySampleAntGroup.Ant_2_initmode_room_0())
		//myAntGroup.Print()

	})
}
