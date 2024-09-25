package room

const (
	_Act_UpdateOneLeave   = "Act.UpdateOneLeave."
	_Act_UpdateOneLeave10 = _Act_UpdateOneLeave + "no one is here!"
)

// =====================================================
func (act *rAction) UpdateOneLeave() []errT {
	room := act.room
	if room.Get.UsedSeats() == 0 {
		//Log.
		return logger.Err.Rlog(Act_UpdateOneLeave_10, nil)
	}
	room.set.usedSeats(room.Get.UsedSeats() - 1)
	return nil
}

// =====================================================
