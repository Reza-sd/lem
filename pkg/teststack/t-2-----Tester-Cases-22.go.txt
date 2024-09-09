package teststack

var Method2_test = AllCasesFunc{
	FuncName: "MyTester.Method2",
	//Skip: true,
	TestCases: []Case{
		//-------------------------------------

		{
			//Skip: true,
			Des: "test int description",
			Set: func() (string, any, any) {
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
