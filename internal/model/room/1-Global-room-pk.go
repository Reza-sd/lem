package room

const (
	PackageName        string = "room"
	MaxSeatsStartEnd   mtr    = 5000
	UsedSeatsStartEnd  mtr    = 1000
	AllSeatsNormalRoom mtr    = 1
	startRoomName      mtr    = 1 //always 1 (use 0 as null)

	maxLenConnectionSlice = 5

)

const ( //error or status codes
	null          statusCode = 0
	sliceOverFlow statusCode = iota + 1
	emptySlice
	exceedCapacity
)

// --------------------
type mtr = uint16
type statusCode uint8

type answer[T any] struct {
	ans   T
	sCode statusCode //I have a problem, here is the problem's code = return specific status code=>0 = nil
	sMsg  string     //"" status massage
	//wCode []statusCode //War sCode
	//wMsg []string //Wrap sMsg
}

//--------------------
