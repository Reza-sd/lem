package room

// ===============================================================
func (act *rAction) UpdateOneCome() statArrT {
	room := act.room
	if room.Get.UsedSeats() == room.Get.AllSeats() || room.Get.UsedSeats()+1 > room.Get.AllSeats() {
		return stat(Act_UpdateOneCome_OverCap, nil)
	}

	room.set.usedSeats(room.Get.UsedSeats() + 1)
	return nil
}

// =====================================================
