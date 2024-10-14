package room

// =================================================
func (act *roomTactT) UpdateOneCome() []errT {

	if act.room.Get.UsedSeats() == act.room.Get.AllSeats() || act.room.Get.UsedSeats()+1 > act.room.Get.AllSeats() {
		return logger.Act.Err.Rlog(_roomTactT_UpdateOneCome_10, nil)
	}

	act.room.set.usedSeats(act.room.Get.UsedSeats() + 1)
	return nil

}

//===================================================
