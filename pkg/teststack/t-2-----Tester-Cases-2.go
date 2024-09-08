package teststack

var Method2_test = TestCasesforFunc{
	FuncName: "MyTester.Method2",
	//Skip: true,
	TestCases: []TestCase{
		//-------------------------------------

		{
			//Skip: true,
			Des: "test int description",
			Case: func() (string, any, any) {
				a := 1
				b := 5
				input := Inp(a, b)
				got := (b - a)
				exp := 4
				return input, got, exp
			},
		},
		//-------------------------------------

	}}
