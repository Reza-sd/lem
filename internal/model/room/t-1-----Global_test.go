package room

import (
	//"fmt"
	teststack "main/pkg/teststack"
	"testing"
)

var (
	MyTester = teststack.Tester{
		AllTestCasesPkStSlice: []AllTestCasesPkSt{
			{TestCasesFunc: Init_Cases},
			{TestCasesFunc: HasOneFreeSeat_Cases},
			{TestCasesFunc: OneComeUpdateSeats_Cases},
			{TestCasesFunc: Print_Cases, Skip: true},
			{TestCasesFunc: OneLeaveUpdateSeats_Cases},
		},
	}
)

// ========================================
func Test_room_pk(t *testing.T) {
	errRunAll := MyTester.RunAll(t)
	if errRunAll != nil {
		t.Errorf("\n\n>>>>>>###### panic RunAll ########<<<<\n%v\n\n", errRunAll)
	}
}

// ========================================
