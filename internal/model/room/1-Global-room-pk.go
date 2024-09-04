package room

const (
	PackageName        = "room"
	MaxSeatsStartEnd   = 5000
	UsedSeatsStartEnd  = 1000
	AllSeatsNormalRoom = 1
	startRoomName      = 1 //always 1 (use 0 as null)
)

// --------------------
type Mtr = uint16

type answer[T any] struct {
	ans      T
	ok       bool  //do the job completely and perfectly
	statCode uint8 //I have a problem, here is the problem's code = return specific status code
	err      error
}

//--------------------
