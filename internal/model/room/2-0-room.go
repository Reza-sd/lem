package room

// ================================
type room struct {
	name            Mtr
	allSeats        Mtr
	usedSeats       Mtr
	connectionSlice []Mtr
	//IsAvailable bool
	get    getter
	set    setter
	update updater
}
type getter struct {
	room *room
}

type setter struct {
	room *room
}
type updater struct {
	room *room
}

// -----------------------
func newRoom() *room { //Constructor=factory function
	r := &room{}
	r.get = getter{room: r}
	r.set = setter{room: r}
	r.update = updater{room: r}
	return r
}

// ---------------------------------------
