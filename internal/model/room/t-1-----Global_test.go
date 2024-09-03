package room

import (
	//"fmt"
	teststack "main/pkg/teststack"
	"testing"
)

var (
	MyTester = teststack.Tester{
		AllTestCasesPkStSlice: []AllTestCasesPkSt{
			//{TestCasesFunc: Init_Cases},
			//{TestCasesFunc: HasOneFreeSeat_Cases},
			//{TestCasesFunc: OneComeUpdateSeats_Cases},
			{TestCasesFunc: Print_Cases},
			//{TestCasesFunc: OneLeaveUpdateSeats_Cases},
			//{TestCasesFunc: OneRandomNextRoom_Cases},
		},
	}
)

// ========================================
func Test_room_pk(t *testing.T) {
	MyTester.RunAll(t)

}

// ========================================
