package room

// =======================================================
type room struct {
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
