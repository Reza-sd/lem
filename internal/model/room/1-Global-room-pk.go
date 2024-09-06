package room

// =========================================
type RT = uint16
type Err = uint8 //error type

// -----------------------------------------
const (
	PkgName            string = "room"
	MaxSeatsStartEnd   RT     = 5000
	UsedSeatsStartEnd  RT     = 1000
	AllSeatsNormalRoom RT     = 1
	startRoomName      RT     = 1 //always 1 (use 0 as null)

	maxLenConnectionSlice = 5
)

// -----------------------------------------
const ( //func or method status code
	Null Err = iota
	GetOneRandomNextRoom
	GetOneRandomNextRoom10
	GetOneRandomNextRoom20

	UpdateOneCome
	UpdateOneCome10

	UpdateOneLeave
	UpdateOneLeave10
)

// var errorMap = map[int]string{ // for log purpose

// 	ErrorCodeUnauthorized: "Unauthorized",
// 	ErrorCodeNotFound: "Not Found",
// 	ErrorCodeInternalServerError: "Internal Server Error",
// 	// ... other error codes

//   }
//=========================================
