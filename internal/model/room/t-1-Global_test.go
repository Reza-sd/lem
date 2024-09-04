package room

import (
	ts "main/pkg/teststack"
	"testing"
)

var (
	MyTester = ts.Tester{
		AllTestCasesPkStSlice: []ts.AllTestCasesPkSt{
			{TestCasesforFunc: test_roomBuilder_Cases},
			{TestCasesforFunc: hasOneFreeSeat_Cases},
			{TestCasesforFunc: oneCome_Cases},
			{TestCasesforFunc: Print_Cases},
			{TestCasesforFunc: oneLeave_Cases},
			{TestCasesforFunc: test_oneRandomNextRoom_Cases},
		},
	}
)

// ========================================
func Test_room_pk(t *testing.T) {
	MyTester.RunAll(t)

}

// ========================================
