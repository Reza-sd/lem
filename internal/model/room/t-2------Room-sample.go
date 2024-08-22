package room

// ====================================
type TestCase struct {
	Skip bool
	Des  string
	got  any
	exp  any
}
type TestCasesFunc struct {
	Skip      bool
	FuncName  string
	TestCases []TestCase
}

var AllTestCasesPkSlice = []TestCasesFunc{Init_Test}

// ====================================
// --------------------------------------
var Sample_Room_End = &Room{
	Name:            1,
	AllSeats:        5000,
	UsedSeats:       1000,
	ConnectionSlice: []Mtr{2, 3, 4},
}

// --------------------------------------
var Sample_Room_Start_name_0 = &Room{
	Name:            0,
	AllSeats:        5000,
	UsedSeats:       1000,
	ConnectionSlice: []Mtr{1, 2, 3},
}

// --------------------------------------
var Sample_Room_Middle = &Room{
	Name:            3,
	AllSeats:        1,
	UsedSeats:       0,
	ConnectionSlice: []Mtr{2, 4},
}

// --------------------------------------
