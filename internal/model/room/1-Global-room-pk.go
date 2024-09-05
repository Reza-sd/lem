package room

const (
	PkgName            string = "room"
	MaxSeatsStartEnd   mtr    = 5000
	UsedSeatsStartEnd  mtr    = 1000
	AllSeatsNormalRoom mtr    = 1
	startRoomName      mtr    = 1 //always 1 (use 0 as null)

	maxLenConnectionSlice = 5
)

// const ( //error or status codes
// 	null          uint8 = 0 //[0-10] reserved
// 	sliceOverFlow uint8 = iota + 10
// 	emptySlice
// 	exceedCapacity
// )

// error is for code
// log is for coder
// error holder stack = map[uint8]uint8
var ErrorsHolder = make(map[uint8]uint8)

const ( //func or method status code
	//-----getter--------------------
	null              = 0
	OneRandomNextRoom = iota + 10
	OneRandomNextRoom_code_10
	OneRandomNextRoom_Code_20

	oneCome
	oneCome_code_10

	oneLeave
	oneLeave_code_10

	//-----setter--------------------

	//--------------------------------
)

// --------------------
type mtr = uint16

//type statusCode uint8

// type answer[T any] struct {
// 	Ans T
// 	//Stat Status
// 	// Err  error
// 	Code uint8 //statusCode
// 	//Msg  string //"" status massage

// 	//wCode []statusCode //War sCode
// 	//wMsg []string //Wrap sMsg
// 	//l logger
// 	//errGen func()

// }

// func err() answer[any]{

// }
// -------------------
