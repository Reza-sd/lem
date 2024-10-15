package room

// ==========================================
//const _EMPTY_USED_SEATS = 0

func (act *roomTactT) UpdateOneLeave() []errT {

	if act.room.Get.UsedSeats() == _EMPTY_USED_SEATS {
		return logger.Act.Err.Rlog(_ERR_roomT_actT_UpdateOneLeave_10, nil, "if room.Get.UsedSeats() == _EMPTY_USED_SEATS")
	}
	act.room.set.usedSeats(act.room.Get.UsedSeats() - 1)
	return nil
}

//==============================================
