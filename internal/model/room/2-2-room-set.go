package room

// ================================
func (s *setter) name(name mtr) *setter {
	s.room.name = name
	return s
}
func (s *setter) allSeats(allSeats mtr) *setter {
	s.room.allSeats = allSeats
	return s
}
func (s *setter) usedSeats(usedSeats mtr) *setter {
	s.room.usedSeats = usedSeats
	return s
}
func (s *setter) connectionSlice(connectionSlice []mtr) *setter {
	s.room.connectionSlice = make([]mtr, len(connectionSlice), maxLenConnectionSlice)
	copy(s.room.connectionSlice, connectionSlice)

	return s
}

// ===============================================================
func (s *setter) oneCome() *room {

	if s.room.get.usedSeats() == s.room.get.allSeats() || s.room.get.usedSeats()+1 > s.room.get.allSeats() {
		return Answer[*room](s.room, Room_set_oneCome, Room_set_oneCome_code_10, s.room)
	}

	s.room.set.usedSeats(s.room.get.usedSeats() + 1)
	return Answer[*room](s.room, Room_set_oneCome, null, s.room)
}

// =====================================================
func (s *setter) oneLeave() *room {
	if s.room.get.usedSeats() == 0 {
		return Answer[*room](s.room, Room_set_oneLeave, Room_set_oneLeave_code_10, s.room)
	}
	s.room.set.usedSeats(s.room.get.usedSeats() - 1)
	return Answer[*room](s.room, Room_set_oneLeave, null, s.room)
}

// =====================================================
