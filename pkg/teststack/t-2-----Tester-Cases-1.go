package teststack

import "fmt"

var myTester1_test = TestCasesforFunc{
	FuncName: "MyTester.Method2",
	//Skip: true,
	TestCases: []TestCase{
		{
			Des: "test1111",
			Setup: func()(any,any,any){
				//Des:="test1111"
				input :="input value"
				got :=3
				exp:=3
				return input,got,exp
			},
			// Got: func() any {
			// 	// a := 1
			// 	// b := 3
			// 	// c := b - a
			// 	return 3
			// },
			Exp: 30,
		},
		{
			//Skip: true,
			Des: "test 22222",
			Setup: func()(any,string){
				return 4,"input value"
			},
			//Got: 4,
			Exp: 4,
		},
		{
			//Skip: true,
			Des: "test 3333",
			Setup: func()(any,string){
				return "mio","input value"
			},
			//Got: "mio",
			Exp: "mio",
		},
		{
			//Skip: true,
			Des: "test 444",
			Setup: func()(any,string){
				return 5,"input value"
			},
			//Got: 5,
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
			Setup: func()(any,string){
				return true,"input value"
			},
			//Got: true,
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
			Setup: func()(any,string){
				return struct{ mio bool }{mio: true},"input value"
			},
			//Got: struct{ mio bool }{mio: true},
			Exp: struct{ mio bool }{mio: true},
		},

		//--------------------------
		{
			//Skip: true,
			Des: "test 77777",
			Setup: func()(any,string){
				return &struct{ mio bool }{mio: true},"input value"
			},
			//Got: &struct{ mio bool }{mio: true},
			Exp: &struct{ mio bool }{mio: true},
		},
		//--------------------------
		{
			//Skip: true,
			Des: "test 8888",
			Setup: func()(any,string){
				return fmt.Errorf("err msg"),"input value"
			},
			//Got: fmt.Errorf("err msg"),
			Exp: fmt.Errorf("err msg"),
		},
		//--------------------------

	}}
