package room

// ===============================================================
func (a *rAction) UpdateOneCome() statArrT {
	room := a.room
	if room.Get.UsedSeats() == room.Get.AllSeats() || room.Get.UsedSeats()+1 > room.Get.AllSeats() {
		return stat(UpdateOneCome10, nil)
	}

	room.Set.UsedSeats(room.Get.UsedSeats() + 1)
	return nil
}

// =====================================================
func (a *rAction) UpdateOneLeave() statArrT {
	room := a.room
	if room.Get.UsedSeats() == 0 {
		return stat(UpdateOneLeave10, nil)
	}
	room.Set.UsedSeats(room.Get.UsedSeats() - 1)
	return nil
}

// =====================================================

func (act *rAction) Print() {

	println()

	print("Room: Name=")
	print(act.room.Get.Name())

	print(", AllSeats=")
	print(act.room.Get.AllSeats())

	print(", UsedSeats=")
	print(act.room.Get.UsedSeats())

	print(", ConnectionSlice[]index:(value)=")
	for index, value := range act.room.Get.ConnectionSlice() {
		print(index, ":(", value, "),")

	}
	println()

}

//=============================================================
