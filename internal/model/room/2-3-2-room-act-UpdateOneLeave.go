package room

const (
	_UpdateOneLeave   = "Act.UpdateOneLeave."
	_UpdateOneLeave10 = "no one is here!"
)

// =====================================================
func (act *rAction) UpdateOneLeave() errArrT {
	room := act.room
	if room.Get.UsedSeats() == 0 {
		//Log.
		return stat(UpdateOneLeave10, nil)
	}
	room.set.usedSeats(room.Get.UsedSeats() - 1)
	return nil
}

// =====================================================
