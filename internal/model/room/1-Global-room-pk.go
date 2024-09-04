package room

const (
	PackageName        = "room"
	MaxSeatsStartEnd   = 5000
	UsedSeatsStartEnd  = 1000
	AllSeatsNormalRoom = 1
	startRoomName      = 1 //always 1 (use 0 as null)

)

const (
	null                     = 0
	SliceOverFlow statusCode = iota + 1
	EmptySlice
)

// --------------------
type Mtr = uint16
type statusCode uint8

type answer[T any] struct {
	ans   T
	sCode statusCode //I have a problem, here is the problem's code = return specific status code=>0 = nil
	sMsg  string     //"" status massage
}

//--------------------
