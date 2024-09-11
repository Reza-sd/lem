package room

const (
	_UpdateOneCome   = "Act.UpdateOneCome."
	_UpdateOneCome10 = "no more free seat"
)

// ===============================================================
func (act *rAction) UpdateOneCome() statArrT {
	room := act.room
	if room.Get.UsedSeats() == room.Get.AllSeats() || room.Get.UsedSeats()+1 > room.Get.AllSeats() {
		return stat(UpdateOneCome10, nil)
	}

	room.set.usedSeats(room.Get.UsedSeats() + 1)
	return nil
}

// =====================================================
