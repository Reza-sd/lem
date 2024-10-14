package room

// ==========================================
func (act *roomTactT) UpdateOneLeave() []errT {

	if act.room.Get.UsedSeats() == 0 {
		return logger.Act.Err.Rlog(_roomTactT_UpdateOneLeave_10, nil, "if room.Get.UsedSeats() == 0")
	}
	act.room.set.usedSeats(act.room.Get.UsedSeats() - 1)
	return nil
}

//==============================================
