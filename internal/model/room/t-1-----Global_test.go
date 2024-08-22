package room

import (
	"reflect"
	"testing"
)

// ====================================
type TestCase struct {
	Des string
	got any
	exp any
}
type TestCasesFunc struct {
	FuncName  string
	TestCases []TestCase
}
type TestCases map[string][]TestCase

var AllTestCasesPkSlice = []TestCasesFunc{Init_Test}

// var  TestCases struct {
// 	map[string][]TestCase{}
// }
//var b= a+2
//====================================
func Test_InitGraph(t *testing.T) {
	for i := 0; i < len(AllTestCasesPkSlice); i++ {
        //fmt.Println(mySlice[i])

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
