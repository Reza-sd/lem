package room

const (
	PackageName        = "room"
	MaxSeatsStartEnd   = 5000
	UsedSeatsStartEnd  = 1000
	AllSeatsNormalRoom = 1
	startRoomName      = 1 //always 1
)

// --------------------
type Mtr = uint16

type answer[T any] struct {
	ans T
	err error
}

//--------------------
