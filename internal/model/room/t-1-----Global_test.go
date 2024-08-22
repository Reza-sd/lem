package room

import (
	teststack "main/pkg/teststack"
	"testing"
)

var (
	MyTester = teststack.Tester{
		//PackageName: "mio",

		AllTestCasesPkStSlice: []AllTestCasesPkSt{
			//{TestCasesFunc: Init_Test, Skip: true},
			{TestCasesFunc: Init_Cases},
			{TestCasesFunc: HasOneFreeSeat_Cases},
			{TestCasesFunc: OneComeUpdateSeats_Cases},
			{TestCasesFunc: Print_Cases, Skip: true},
		},
	}
)

// ========================================
func Test_room_pk(t *testing.T) {
	MyTester.RunAll(t)
	//MyTester.RunOne(t, Init_Test, true)
	//MyTester.RunOne(t, Init_Test, false)
}

// ========================================
