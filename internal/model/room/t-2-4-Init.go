package room

var Init_Test = TestCasesFunc{
	FuncName: "Init",
	TestCases: []TestCase{
		{
			Des: "test1111",
			got: func() any {
				a := 1
				b := 3
				c := b - a
				return Mtr(c)
			},
			exp: Mtr(3),
		},
		{
			Des: "test 22222",
			got: 4,
			exp: 4,
		},
	}}
