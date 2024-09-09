package room

// ===============================================================
func (a *roomAction) UpdateOneCome() (*roomAction, []statType) {
	room := a.room
	if room.get.UsedSeats() == room.get.AllSeats() || room.get.UsedSeats()+1 > room.get.AllSeats() {
		return a, wrapper(UpdateOneCome10, nil)
	}

	room.set.UsedSeats(room.get.UsedSeats() + 1)
	return a, nil
}

// =====================================================
func (a *roomAction) UpdateOneLeave() (*roomAction, []statType) {
	room := a.room
	if room.get.UsedSeats() == 0 {
		return a, wrapper(UpdateOneLeave10, nil)
	}
	room.set.UsedSeats(room.get.UsedSeats() - 1)
	return a, nil
}

// =====================================================
