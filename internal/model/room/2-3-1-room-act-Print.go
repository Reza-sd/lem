package room

// =============================================================
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
