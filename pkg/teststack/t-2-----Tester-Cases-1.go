package teststack

var Method1_test = TestCasesforFunc{
	FuncName: "MyTester.Method2",
	//Skip: true,
	TestCases: []TestCase{
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
		//--------------------------
		{
			//Skip: true,
			Des: "test string",
			Case: func() (string, any, any) {
				inp1 := "ali"
				inp2 := "reza"
				input := Inp(inp1)
				got := inp1 + inp2
				exp := "alireza"
				return input, got, exp
			},
		},

		//--------------------------
		{
			//Skip: true,
			Des: "test struct",
			Case: func() (string, any, any) {
				inp1 := true
				input := Inp(inp1)
				got := struct{ mio bool }{mio: inp1}
				exp := struct{ mio bool }{mio: true}
				return input, got, exp
			},
		},
		//--------------------------
		{
			//Skip: true,
			Des: "test error type",
			Case: func() (string, any, any) {
				var inp1 error
				//Des:="test1111"
				input := Inp(inp1)
				got := inp1
				var out error
				exp := out
				return input, got, exp
			},
		},
		//--------------------------

	}}
