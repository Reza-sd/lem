package room

import (
	"fmt"
	"reflect"
	"testing"
)

// ========================================
func Test_RoomPk(t *testing.T) {
	for i := 0; i < len(AllTestCasesPkSlice); i++ {
		//fmt.Println(mySlice[i])
		for j := 0; j < len(AllTestCasesPkSlice[i].TestCases); j++ {

			NumfuncDes := fmt.Sprintf("%v-%v-%v", j+1, AllTestCasesPkSlice[i].FuncName, AllTestCasesPkSlice[i].TestCases[j].Des)

			t.Run(NumfuncDes, func(t *testing.T) {
				if AllTestCasesPkSlice[i].TestCases[j].Skip {
					t.Skip()
				}
				got := AllTestCasesPkSlice[i].TestCases[j].got
				//print(got)
				//a :=*got
				//fmt.Println("***>",got)
				// switch got.(type){
				// case func()any:
				// 	fmt.Println("BBBBBBBB")
				// }
				exp := AllTestCasesPkSlice[i].TestCases[j].exp
				Assert(t, got, exp)
			})
		}

	}
}

// ==================================
func Assert(t testing.TB, got, exp any) {
	t.Helper()

	var gotValue any
	var expValue any

	switch got.(type) {
	case func() any:
		gotValue = got.(func() any)()
	default:
		gotValue = got
	}

	switch exp.(type) {
	case func() any:
		expValue = exp.(func() any)()
	default:
		expValue = exp
	}

	if !reflect.DeepEqual(gotValue, expValue) {
		//t.Fatalf("\n>>>not same: \n got=%v<<\n exp=%v<<", got, exp)
		t.Errorf("\n>>>not same: \n got=%v<<\n exp=%v<<", got, exp)
	}
}

//======================================
