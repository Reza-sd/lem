package room

import (
	"fmt"
	"reflect"
	"testing"
)

// ========================================
func Test_InitGraph(t *testing.T) {
	for i := 0; i < len(AllTestCasesPkSlice); i++ {
		//fmt.Println(mySlice[i])
		for j := 0; j < len(AllTestCasesPkSlice[i].TestCases); j++ {
			NumfuncDes := fmt.Sprintf("%v-%v-%v", j+1, AllTestCasesPkSlice[i].FuncName, AllTestCasesPkSlice[i].TestCases[j].Des)

			t.Run(NumfuncDes, func(t *testing.T) {

				got := AllTestCasesPkSlice[i].TestCases[j].got
				//a :=*got
				exp := AllTestCasesPkSlice[i].TestCases[j].exp
				Assert(t, got, exp)
			})
		}

	}
}

// ==================================
func Assert(t testing.TB, got, exp any) {
	t.Helper()
	if !reflect.DeepEqual(got, exp) {
		//t.Fatalf("\n>>>not same: \n got=%v<<\n exp=%v<<", got, exp)
		t.Errorf("\n>>>not same: \n got=%v<<\n exp=%v<<", got, exp)
		//print(TT.Des)
	}
}

//======================================
