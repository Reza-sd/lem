package room

import (
	//"reflect"
	"testing"
)

func Test_Init(t *testing.T) {
	//---------------------------------------------
	t.Run("1-Init if name=0 means start home", func(t *testing.T) {
		//sample := SampleRoom{}
		//myRoom := sample.Sample_Room_1_5000_1000()
		myRoom := Room{}
		myRoom.Init(0, 5, []Mtr{1, 2, 3})
		expRoom := Room{
			Name:            0,
			AllSeats:        5000,
			UsedSeats:       1000,
			ConnectionSlice: []Mtr{1, 2, 3},
		}
		Assert(t, myRoom, expRoom)
		//myRoom.Print()
	})
	//---------------------------------------------
	t.Run("1-Init non start or end room", func(t *testing.T) {
		//sample := SampleRoom{}
		//myRoom := sample.Sample_Room_1_5000_1000()
		myRoom := Room{}
		myRoom.Init(1, 5, []Mtr{1, 2, 3})
		expRoom := Room{
			Name:            1,
			AllSeats:        1,
			UsedSeats:       0,
			ConnectionSlice: []Mtr{1, 2, 3},
		}
		Assert(t, myRoom, expRoom)
		//myRoom.Print()
	})
	//---------------------------------------------
	t.Run("1-Init end room", func(t *testing.T) {
		//sample := SampleRoom{}
		//myRoom := sample.Sample_Room_1_5000_1000()
		myRoom := Room{}
		myRoom.Init(5, 5, []Mtr{1, 2, 3})
		expRoom := Room{
			Name:            5,
			AllSeats:        5000,
			UsedSeats:       1000,
			ConnectionSlice: []Mtr{1, 2, 3},
		}
		Assert(t, myRoom, expRoom)
		//myRoom.Print()
	})
}

// ==================================
// func assert_Init(t testing.TB, myRoom, expRoom Room) {
// 	t.Helper()
// 	if !reflect.DeepEqual(myRoom, expRoom) {
// 		t.Errorf("\n>>>not same: \n myRoom=%v<<\n expRoom=%v<<", myRoom, expRoom)
// 	}
// }

//======================================
