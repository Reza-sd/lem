package room

import (
	"fmt"
	"reflect"
	"testing"
)

// ====================================

type SampleRoomSt struct{}

var SampleRoom = SampleRoomSt{}

// ====================================
func (s *SampleRoomSt) End_Name_1() *room {
	r := &room{
		data: data{
			name:            1,
			allSeats:        MaxSeatsStartEnd,
			usedSeats:       UsedSeatsStartEnd,
			connectionSlice: []RT{},
		},
	}
	r.Get.room = r
	r.set.room = r
	r.Act.room = r
	return r
}

// --------------------------------------
func (s *SampleRoomSt) Start_Name_0() *room {
	r := &room{
		data: data{
			name:            0,
			allSeats:        MaxSeatsStartEnd,
			usedSeats:       UsedSeatsStartEnd,
			connectionSlice: []RT{},
		},
	}
	r.Get.room = r
	r.set.room = r
	r.Act.room = r
	return r

}

// --------------------------------------
func (s *SampleRoomSt) Middle_Name_3() *room {
	r := &room{
		data: data{
			name:            3,
			allSeats:        1,
			usedSeats:       0,
			connectionSlice: []RT{},
		},
	}
	r.Get.room = r
	r.set.room = r
	r.Act.room = r
	return r

}

// ==========================================================
// func assert2[T any](t testing.TB, got, exp T, inputsArr ...any) {
// 	t.Helper()
// 	inputsStr := Inp(inputsArr)
// 	if !reflect.DeepEqual(got, exp) {
// 		t.Errorf("\n>------------------\nInputs:\n%v\nGot=type:(%T) value:(%v)\nExp=type:(%T) value:(%v)\n>------------------\n",
// 			inputsStr, got, got, exp, exp)
// 	}
// }

// ==========================================================
func assert(t testing.TB, got, exp any, inputsArr ...any) {
	t.Helper()
	inputsStr := InputsArrToString(inputsArr)

	if !reflect.DeepEqual(got, exp) {
		t.Errorf("\n>------------------\nInputs:\n%v\nGot=type:(%T) value:(%v)\nExp=type:(%T) value:(%v)\n>------------------\n",
			inputsStr, got, got, exp, exp)
	}
}

// ==========================================================
func InputsArrToString(inputs []any) string {
	var str string
	var count int
	for _, item := range inputs {
		count++
		str = str + fmt.Sprintf("%d-inp=type:(%T) value:(%v)\n", count, item, item)
	}
	return str
}

// ==========================================================
