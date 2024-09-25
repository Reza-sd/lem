package room

// =======================================================
type room struct {
	data data //private

	//categorise behaviour
	set rSetter //private

	Get rGetter //<--export (public)
	Act rAction //<--export (public)
}
type data struct {
	name            RT
	allSeats        RT
	usedSeats       RT
	connectionSlice []RT
}

//==========================================================
