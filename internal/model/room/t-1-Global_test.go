package room

import (
	//"fmt"
	teststack "main/pkg/teststack"
	"testing"
)

var (
	MyTester = teststack.Tester{
		AllTestCasesPkStSlice: []AllTestCasesPkSt{
			{TestCasesFunc: builder_Cases},
			{TestCasesFunc: hasOneFreeSeat_Cases},
			{TestCasesFunc: oneCome_Cases},
			{TestCasesFunc: Print_Cases},
			{TestCasesFunc: oneLeave_Cases},
			{TestCasesFunc: OneRandomNextRoom_Cases},
		},
	}
)

// ========================================
func Test_room_pk(t *testing.T) {
	MyTester.RunAll(t)

}

// ========================================
