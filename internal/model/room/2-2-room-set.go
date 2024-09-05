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
func (s *setter) oneCome() answer[*room] {

	if s.room.get.usedSeats() == s.room.get.allSeats() || s.room.get.usedSeats()+1 > s.room.get.allSeats() {
		return answer[*room]{Ans: s.room, Code: exceedCapacity, Msg: "no more seats"}
	}

	s.room.set.usedSeats(s.room.get.usedSeats() + 1)
	return answer[*room]{Ans: s.room}
}

// =====================================================
func (s *setter) oneLeave() answer[*room] {
	if s.room.get.usedSeats() == 0 {
		return answer[*room]{Ans: s.room, Code: emptySlice, Msg: "seats empty"}
	}
	s.room.set.usedSeats(s.room.get.usedSeats() - 1)
	return answer[*room]{Ans: s.room}
}

// =====================================================
