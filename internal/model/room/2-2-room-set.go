package room

// ================================
func (s *roomSetter) Name(name RT) statArrTyp {
	//check if name valid to set
	s.room.data.name = name
	return nil
}
func (s *roomSetter) AllSeats(allSeats RT) statArrTyp {
	//check validation
	s.room.data.allSeats = allSeats
	return nil
}
func (s *roomSetter) UsedSeats(usedSeats RT) statArrTyp {
	s.room.data.usedSeats = usedSeats
	return nil
}
func (s *roomSetter) ConnectionSlice(connectionSlice []RT) statArrTyp {
	s.room.data.connectionSlice = make([]RT, len(connectionSlice), maxLenConnectionSlice)
	copy(s.room.data.connectionSlice, connectionSlice)

	return nil
}

// ================================
