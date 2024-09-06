package room

// =========================================
type m = uint16
type e = uint8 //error type

// -----------------------------------------
const (
	PkgName            string = "room"
	MaxSeatsStartEnd   m      = 5000
	UsedSeatsStartEnd  m      = 1000
	AllSeatsNormalRoom m      = 1
	startRoomName      m      = 1 //always 1 (use 0 as null)

	maxLenConnectionSlice = 5
)

// -----------------------------------------
const ( //func or method status code
	Null e = iota
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
