package teststack

var myTester_Cases_1 = TestCasesforFunc{
	FuncName: "MyTester.Method1",
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
			Des: "test 444",
			Got: 5,
			Exp: func() any {
				a := 1
				b := 6
				c := b - a
				return c
			},
		},
		//--------------------------
		{
			//Skip: true,
			Des: "test 5555",
			Got: true,
			Exp: func() any {
				a := 1
				b := 6
				c := b - a
				return c == 5
			},
		},
		//--------------------------
		{
			//Skip: true,
			Des: "test 6666",
			Got: struct{ mio bool }{mio: true},
			Exp: struct{ mio bool }{mio: true},
		},

		//--------------------------
		{
			//Skip: true,
			Des: "test 77777",
			Got: &struct{ mio bool }{mio: true},
			Exp: &struct{ mio bool }{mio: true},
		},

		//--------------------------
		{
			//Skip: true,
			Des: "test 8888",
			Got: func() any {
				a := 1
				b := 6
				c := b - a
				return c
			},
			Exp: func() any {
				a := 2
				b := 7
				c := b - a
				return c
			},
		},
		//--------------------------

		{
			//Skip: true,
			Des: "test 9999",
			Got: func() any {

				return struct{ mio bool }{mio: true}
			},
			Exp: func() any {

				return struct{ mio bool }{mio: true}
			},
		},
		//--------------------------
	}}
