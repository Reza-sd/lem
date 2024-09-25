package room

const (
	_UpdateOneCome   = "Act.UpdateOneCome."
	_UpdateOneCome10 = _UpdateOneCome + "no more free seat"
)

// ===============================================================
func (act *rAction) UpdateOneCome() errArrT {
	room := act.room
	if room.Get.UsedSeats() == room.Get.AllSeats() || room.Get.UsedSeats()+1 > room.Get.AllSeats() {
		return logger.Err.Rlog(UpdateOneCome10, nil)
	}

	room.set.usedSeats(room.Get.UsedSeats() + 1)
	return nil

}

// =====================================================
