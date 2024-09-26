package room

// =======================================================
// Struct in Golang is a user-defined data type
type roomT struct {
	data roomTdataT //private

	//categorise behaviour
	set roomTsetT //private

	Get roomTgetT //<--export (public)
	Act roomTactT //<--export (public)
}
type roomTdataT struct {
	name            rT
	allSeats        rT
	usedSeats       rT
	connectionSlice []rT
}

//==========================================================
