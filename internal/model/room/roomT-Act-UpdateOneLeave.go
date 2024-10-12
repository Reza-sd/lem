package room

// ==========================================
func (act *roomTactT) UpdateOneLeave() []errT {
	room := act.room
	if room.Get.UsedSeats() == 0 {
		return logger.Act.Err.Rlog(roomTactT_UpdateOneLeave_10, nil)
	}
	room.set.usedSeats(room.Get.UsedSeats() - 1)
	return nil
}

//==============================================
