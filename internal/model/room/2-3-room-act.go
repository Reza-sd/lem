package room

// ===============================================================
func (a *roomAction) UpdateOneCome() statArrTyp {
	room := a.room
	if room.get.UsedSeats() == room.get.AllSeats() || room.get.UsedSeats()+1 > room.get.AllSeats() {
		return wrapper(UpdateOneCome10, nil)
	}

	room.set.UsedSeats(room.get.UsedSeats() + 1)
	return nil
}

// =====================================================
func (a *roomAction) UpdateOneLeave() statArrTyp {
	room := a.room
	if room.get.UsedSeats() == 0 {
		return wrapper(UpdateOneLeave10, nil)
	}
	room.set.UsedSeats(room.get.UsedSeats() - 1)
	return nil
}

// =====================================================

func (act *roomAction) Print() {

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
