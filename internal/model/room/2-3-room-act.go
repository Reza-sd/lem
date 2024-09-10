package room

// ===============================================================
func (a *rAction) UpdateOneCome() statArrT {
	room := a.room
	if room.get.UsedSeats() == room.get.AllSeats() || room.get.UsedSeats()+1 > room.get.AllSeats() {
		return stat(UpdateOneCome10, nil)
	}

	room.set.UsedSeats(room.get.UsedSeats() + 1)
	return nil
}

// =====================================================
func (a *rAction) UpdateOneLeave() statArrT {
	room := a.room
	if room.get.UsedSeats() == 0 {
		return stat(UpdateOneLeave10, nil)
	}
	room.set.UsedSeats(room.get.UsedSeats() - 1)
	return nil
}

// =====================================================

func (act *rAction) Print() {

	println()

	print("Room: Name=")
	print(act.room.get.Name())

	print(", AllSeats=")
	print(act.room.get.AllSeats())

	print(", UsedSeats=")
	print(act.room.get.UsedSeats())

	print(", ConnectionSlice[]index:(value)=")
	for index, value := range act.room.get.ConnectionSlice() {
		print(index, ":(", value, "),")

	}
	println()

}

//=============================================================
