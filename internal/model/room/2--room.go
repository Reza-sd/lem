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

// -----------------------
func newRoom() *room { //Constructor=factory function
	r := &room{}
	r.get = getter{room: r}
	r.set = setter{room: r}
	r.update = updater{room: r}
	return r
}

// ---------------------------------------
type getter struct {
	room *room
}
type setter struct {
	room *room
}
type updater struct {
	room *room
}

// -------Getter---------------
func (get *getter) name() Mtr {
	return get.room.name
}

func (get *getter) allSeats() Mtr {
	return get.room.allSeats
}
func (get *getter) usedSeats() Mtr {
	return get.room.usedSeats
}
func (get *getter) connectionSlice() []Mtr {
	return get.room.connectionSlice
}

// -----------Setter------------------
func (set *setter) name(name Mtr) {
	set.room.name = name
}
func (set *setter) allSeats(allSeats Mtr) {
	set.room.allSeats = allSeats
}
func (set *setter) usedSeats(usedSeats Mtr) {
	set.room.usedSeats = usedSeats
}
func (set *setter) connectionSlice(connectionSlice []Mtr) {
	set.room.connectionSlice = connectionSlice
}

//----------------------------------------
