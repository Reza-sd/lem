package room

// ================================
type Room struct {
	name            Mtr
	allSeats        Mtr
	usedSeats       Mtr
	connectionSlice []Mtr
	//IsAvailable bool
}

// -------Getter---------------
func (r *Room) getName() Mtr {
	return r.name
}
func (r *Room) getAllSeats() Mtr {
	return r.allSeats
}
func (r *Room) getUsedSeats() Mtr {
	return r.usedSeats
}
func (r *Room) getConnectionSlice() []Mtr {
	return r.connectionSlice
}

// -----------Setter------------------
func (r *Room) setName(name Mtr) {
	r.name = name
}
func (r *Room) setAllSeats(allSeats Mtr) {
	r.allSeats = allSeats
}
func (r *Room) setUsedSeats(usedSeats Mtr) {
	r.usedSeats = usedSeats
}
func (r *Room) setConnectionSlice(connectionSlice []Mtr) {
	r.connectionSlice = connectionSlice
}

//----------------------------------------
