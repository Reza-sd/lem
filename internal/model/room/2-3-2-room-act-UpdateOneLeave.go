package room

// =====================================================
func (act *rAction) UpdateOneLeave() statArrT {
	room := act.room
	if room.Get.UsedSeats() == 0 {
		return stat(Act_UpdateOneLeave_OverCap, nil)
	}
	room.set.usedSeats(room.Get.UsedSeats() - 1)
	return nil
}

// =====================================================
