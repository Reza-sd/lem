package room

// ================================
type rSetter struct {
	room *room
}

// ================================
const (
	_set_name    = "set.name."
	_set_name_10 = _set_name + ""
)

func (set *rSetter) name(name RT) errArrT {
	if name > maxName {
		//return logger.Err.Rlog(,nil)
	}
	//check if name valid to set
	set.room.data.name = name
	return nil
}

//-------------------------------------------------
func (set *rSetter) allSeats(allSeats RT) errArrT {
	//check validation
	set.room.data.allSeats = allSeats
	return nil
}
func (set *rSetter) usedSeats(usedSeats RT) errArrT {
	set.room.data.usedSeats = usedSeats
	return nil
}
func (set *rSetter) connectionSlice(connectionSlice []RT) errArrT {
	set.room.data.connectionSlice = make([]RT, len(connectionSlice), maxLenConnectionSlice)
	copy(set.room.data.connectionSlice, connectionSlice)

	return nil
}

//================================
