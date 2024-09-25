package room

// ================================
type rSetter struct {
	room *room
}

// ================================
const (
	_set_name    = "set.name."
	_set_name_10 = _set_name + "10:exceed max Name"
)

func (set *rSetter) name(name RT) []errT {
	//Guard clause
	if name > maxName {
		return logger.Err.Rlog(set_name_10, nil, "name:",name, ">","maxName:",maxName)
	}
	//check if name valid to set
	set.room.data.name = name
	return nil
}

//-------------------------------------------------
func (set *rSetter) allSeats(allSeats RT) []errT {
	//check validation
	set.room.data.allSeats = allSeats
	return nil
}
func (set *rSetter) usedSeats(usedSeats RT) []errT {
	set.room.data.usedSeats = usedSeats
	return nil
}
func (set *rSetter) connectionSlice(connectionSlice []RT) []errT {
	set.room.data.connectionSlice = make([]RT, len(connectionSlice), maxLenConnectionSlice)
	copy(set.room.data.connectionSlice, connectionSlice)

	return nil
}

//================================
