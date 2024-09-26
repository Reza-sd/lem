package room

// =======================================================
//Struct in Golang is a user-defined data type
type roomT struct {
	data dataT //private

	//categorise behaviour
	set setT //private

	Get getT //<--export (public)
	Act actT //<--export (public)
}
type dataT struct {
	name            rT
	allSeats        rT
	usedSeats       rT
	connectionSlice []rT
}

//==========================================================
