package room

var Init_Test = TestCasesFunc{
	FuncName: "Init",
	TestCases: []TestCase{
		{
			Des: "test1111",
			got: func() any {
				// a := 1
				// b := 3
				// c := b - a
				return 3
			},
			exp: 3,
		},
		{
			//Skip: true,
			Des: "test 22222",
			got: 4,
			exp: 4,
		},
		{
			//Skip: true,
			Des: "test 3333",
			got: "mio",
			exp: "mio",
		},
		{
			//Skip: true,
			Des: "test 3333",
			got: 5,
			exp: func() any {
				a := 1
				b := 6
				c := b - a
				return c
			},
		},
		{
			//Skip: true,
			Des: "test 3333",
			got: true,
			exp: func() any {
				a := 1
				b := 6
				c := b - a
				return c == 5
			},
		},
	}}
