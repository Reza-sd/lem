package room

// ================================
func (s *roomSetter) Name(name RT) *roomSetter {
	s.room.data.name = name
	return s
}
func (s *roomSetter) AllSeats(allSeats RT) *roomSetter {
	s.room.data.allSeats = allSeats
	return s
}
func (s *roomSetter) UsedSeats(usedSeats RT) *roomSetter {
	s.room.data.usedSeats = usedSeats
	return s
}
func (s *roomSetter) SetConnectionSlice(connectionSlice []RT) *roomSetter {
	s.room.data.connectionSlice = make([]RT, len(connectionSlice), maxLenConnectionSlice)
	copy(s.room.data.connectionSlice, connectionSlice)

	return s
}

// ================================
