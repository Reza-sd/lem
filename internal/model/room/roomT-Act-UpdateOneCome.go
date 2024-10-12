package room

// =================================================
func (act *roomTactT) UpdateOneCome() []errT {
	room := act.room
	if room.Get.UsedSeats() == room.Get.AllSeats() || room.Get.UsedSeats()+1 > room.Get.AllSeats() {
		return logger.Act.Err.Rlog(roomTactT_UpdateOneCome_10, nil)
	}

	room.set.usedSeats(room.Get.UsedSeats() + 1)
	return nil

}

//===================================================
