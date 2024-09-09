package teststack

var funcReza_Cases = map[string]any{
	"return 3 if...": func()any {
		skip := false
		inp1 := 1
		inp2 := 5
		//inp := Inp(inp1, inp2)
		got := (inp2 - inp1)
		exp := 4
		return []any{skip, got, exp, inp1, inp2}
	},
}

//-------------------
// var Method1_test = AllCasesFunc{
// 	FuncName: "MyTester.Method1",
// 	//Skip: true,
// 	TestCases: []Case{
// 		{
// 			//Skip: true,
// 			Des: "test int description",
// 			Set: func() (string, any, any) {
// 				a := 1
// 				b := 5
// 				inp := Inp(a, b)
// 				got := (b - a)
// 				exp := 4
// 				return inp, got, exp
// 			},
// 		},
// 		//--------------------------
// 		{
// 			//Skip: true,
// 			Des: "test string",
// 			Set: func() (string, any, any) {
// 				inp1 := "ali"
// 				inp2 := "reza"
// 				inp := Inp(inp1)
// 				got := inp1 + inp2
// 				exp := "alireza"
// 				return inp, got, exp
// 			},
// 		},

// 		//--------------------------
// 		{
// 			//Skip: true,
// 			Des: "test struct",
// 			Set: func() (string, any, any) {
// 				inp1 := true
// 				inp := Inp(inp1)
// 				got := struct{ mio bool }{mio: inp1}
// 				exp := struct{ mio bool }{mio: true}
// 				return inp, got, exp
// 			},
// 		},
// 		//--------------------------
// 		{
// 			//Skip: true,
// 			Des: "test error type",
// 			Set: func() (string, any, any) {
// 				var inp1 error
// 				//Des:="test1111"
// 				inp := Inp(inp1)
// 				got := inp1
// 				var out error
// 				exp := out
// 				return inp, got, exp
// 			},
// 		},
// 		//--------------------------

// 	}}
