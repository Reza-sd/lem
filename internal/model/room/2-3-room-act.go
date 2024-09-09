package room

// ===============================================================
func (a *roomAction) UpdateOneCome() (*roomAction, []statType) {
	room := a.room
	if room.get.UsedSeats() == room.get.AllSeats() || room.get.UsedSeats()+1 > room.get.AllSeats() {
		return a, wrapper(UpdateOneCome10, nil)
	}

	room.set.UsedSeats(room.get.UsedSeats() + 1)
	return a, nil
}

// =====================================================
func (a *roomAction) UpdateOneLeave() (*roomAction, []statType) {
	room := a.room
	if room.get.UsedSeats() == 0 {
		return a, wrapper(UpdateOneLeave10, nil)
	}
	room.set.UsedSeats(room.get.UsedSeats() - 1)
	return a, nil
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
	//var a []int
	//print(len(a))
	// println(a==nil)
	// a=[]int{1}
	// println(a[0])
	// clear(a)
	// println(a)
	// println(a==nil)
	// a=nil
	// println(a==nil)
}

//=============================================================
