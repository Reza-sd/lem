package room

const (
	PkgName            string = "room"
	MaxSeatsStartEnd   mtr    = 5000
	UsedSeatsStartEnd  mtr    = 1000
	AllSeatsNormalRoom mtr    = 1
	startRoomName      mtr    = 1 //always 1 (use 0 as null)

	maxLenConnectionSlice = 5
)

const ( //func or method status code
	//-----getter--------------------
	null = iota
	Room_get_OneRandomNextRoom
	Room_get_OneRandomNextRoom_code_10
	Room_get_OneRandomNextRoom_Code_20

	Room_set_oneCome
	Room_set_oneCome_code_10

	Room_UpdateOneLeave
	Room_UpdateOneLeave_code_10

	//-----setter--------------------

	//--------------------------------
)

type mtr = uint16

// =========================================
