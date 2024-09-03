package room

//import "fmt"

// =====================================================
func (u *updater) oneLeave() *updater {
	if u.room.get.usedSeats() == 0 {
		//return fmt.Errorf("full")
		return u
	}
	u.room.set.usedSeats(u.room.get.usedSeats() - 1)
	return u
}

// =====================================================
