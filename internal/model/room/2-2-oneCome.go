package room

//import "fmt"

// =====================================================
func (u *updater) oneCome() *updater {
	//u.room
	if u.room.get.usedSeats() == u.room.get.allSeats() || u.room.get.usedSeats()+1 > u.room.get.allSeats() {
		//return fmt.Errorf("full")
		return u
	}

	u.room.set.usedSeats(u.room.get.usedSeats() + 1)
	return u
}

//=====================================================
