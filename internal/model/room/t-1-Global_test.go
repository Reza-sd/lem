package room

import (
	//"fmt"
	teststack "main/pkg/teststack"
	"testing"
)

var (
	MyTester = teststack.Tester{
		AllTestCasesPkStSlice: []AllTestCasesPkSt{
			{TestCasesFunc: test_roomBuilder_Cases},
			{TestCasesFunc: hasOneFreeSeat_Cases},
			{TestCasesFunc: oneCome_Cases},
			{TestCasesFunc: Print_Cases},
			{TestCasesFunc: oneLeave_Cases},
			{TestCasesFunc: test_oneRandomNextRoom_Cases},
		},
	}
)

// ========================================
func Test_room_pk(t *testing.T) {
	MyTester.RunAll(t)

}

// ========================================
