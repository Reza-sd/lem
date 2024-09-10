package room

// ================================
func (s *rSetter) Name(name RT) statArrT {
	//check if name valid to set
	s.room.data.name = name
	return nil
}
func (s *rSetter) AllSeats(allSeats RT) statArrT {
	//check validation
	s.room.data.allSeats = allSeats
	return nil
}
func (s *rSetter) UsedSeats(usedSeats RT) statArrT {
	s.room.data.usedSeats = usedSeats
	return nil
}
func (s *rSetter) ConnectionSlice(connectionSlice []RT) statArrT {
	s.room.data.connectionSlice = make([]RT, len(connectionSlice), maxLenConnectionSlice)
	copy(s.room.data.connectionSlice, connectionSlice)

	return nil
}

// ================================
