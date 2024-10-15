package room

// =======================================================
// Struct in Golang is a user-defined data type
type roomT struct {
	data struct {
		name            rT
		allSeats        rT
		usedSeats       rT
		connectionSlice []rT
	} //private

	//categorise behaviour
	set roomTsetT //private

	Get roomTgetT //<--export (public)
	Act roomTactT //<--export (public)

}

// Act
type roomTactT struct {
	room *roomT //private
}
type roomTgetT struct {
	room *roomT
}

type roomTsetT struct {
	room *roomT
}

//==========================================================
