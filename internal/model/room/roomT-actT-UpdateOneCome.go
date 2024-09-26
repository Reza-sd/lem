package room

const (
	_Act_UpdateOneCome   = "Act.UpdateOneCome."
	_Act_UpdateOneCome10 = _Act_UpdateOneCome + "10:no more free seat"
)

// ===================================================
func (act *roomTactT) UpdateOneCome() []errT {
	room := act.room
	if room.Get.UsedSeats() == room.Get.AllSeats() || room.Get.UsedSeats()+1 > room.Get.AllSeats() {
		return logger.Err.Rlog(roomT_actT_UpdateOneCome_10, nil)
	}

	room.set.usedSeats(room.Get.UsedSeats() + 1)
	return nil

}

//===================================================
