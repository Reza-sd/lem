package room

var Init_Test = TestCasesFunc{
	FuncName: "Init",
	TestCases: []TestCase{
		{
			Des: "Test basic addition",
			got: func() any { return 1 + 2 },
			exp: 5,
		},
		{
			Des: "Test basic subtraction",
			got: 4,
			exp: 3,
		},
	}}


