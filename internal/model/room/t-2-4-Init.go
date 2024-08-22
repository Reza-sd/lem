package room

var Init_Test = TestCasesFunc{
	FuncName: "Init",
	//Skip: true,
	TestCases: []TestCase{
		{
			Des: "test1111",
			Got: func() any {
				// a := 1
				// b := 3
				// c := b - a
				return 3
			},
			Exp: 3,
		},
		{
			//Skip: true,
			Des: "test 22222",
			Got: 4,
			Exp: 4,
		},
		{
			//Skip: true,
			Des: "test 3333",
			Got: "mio",
			Exp: "mio",
		},
		{
			//Skip: true,
			Des: "test 3333",
			Got: 5,
			Exp: func() any {
				a := 1
				b := 6
				c := b - a
				return c
			},
		},
		{
			//Skip: true,
			Des: "test 3333",
			Got: true,
			Exp: func() any {
				a := 1
				b := 6
				c := b - a
				return c == 5
			},
		},
	}}
