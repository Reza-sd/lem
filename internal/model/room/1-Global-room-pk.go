package room

const (
	PkgName            string = "room"
	MaxSeatsStartEnd   mtr    = 5000
	UsedSeatsStartEnd  mtr    = 1000
	AllSeatsNormalRoom mtr    = 1
	startRoomName      mtr    = 1 //always 1 (use 0 as null)

	maxLenConnectionSlice = 5
)

const ( //error or status codes
	null          uint8 = 0 //[0-10] reserved
	sliceOverFlow uint8 = iota + 10
	emptySlice
	exceedCapacity
)

// --------------------
type mtr = uint16

//type statusCode uint8

type answer[T any] struct {
	ans T
	Stat Status
	//l logger
	//errGen func()

}

// --------------------
type Status struct {
	// Err  error
	Code uint8
	Msg  string
	//sCode statusCode //I have a problem, here is the problem's code = return specific status code=>0 = nil
	//sMsg  string     //"" status massage
	//wCode []statusCode //War sCode
	//wMsg []string //Wrap sMsg
}
