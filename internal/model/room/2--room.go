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

// --------------------
type answer[T any] struct {
	ans T
	err error
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

type setter struct {
	room *room
}
type updater struct {
	room *room
}

// -----------Setter------------------
func (set *setter) name(name Mtr) *setter {
	set.room.name = name
	return set
}
func (set *setter) allSeats(allSeats Mtr) *setter {
	set.room.allSeats = allSeats
	return set
}
func (set *setter) usedSeats(usedSeats Mtr) *setter {
	set.room.usedSeats = usedSeats
	return set
}
func (set *setter) connectionSlice(connectionSlice []Mtr) *setter {
	set.room.connectionSlice = connectionSlice
	return set
}

//----------------------------------------
