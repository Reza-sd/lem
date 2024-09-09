package room

import (
	ts "main/pkg/teststack"
	"testing"
)

var (
	MyTester = ts.Tester{
		AllTestCasesPkStSlice: []ts.AllCasesPkg{
			{CasesMap: test_roomBuilder_Cases},
			{CasesMap: hasOneFreeSeat_Cases},
			{CasesMap: oneCome_Cases},
			{CasesMap: Print_Cases},
			{CasesMap: oneLeave_Cases},
			{CasesMap: test_oneRandomNextRoom_Cases},
		},
	}
)

// ========================================
func Test_room_pk(t *testing.T) {
	MyTester.RunAll(t)

}

// ========================================
