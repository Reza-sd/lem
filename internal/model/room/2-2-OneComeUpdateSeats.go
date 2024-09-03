package room

import "fmt"

// =====================================================
func (u *updater) oneCome() error {
	//u.room
	if u.room.get.usedSeats() == u.room.get.allSeats() {
		return fmt.Errorf("full")
	}

	if u.room.get.usedSeats()+1 > u.room.get.allSeats() {
		return fmt.Errorf("full")
	}

	//r.usedSeats++
	u.room.set.usedSeats(u.room.get.usedSeats() + 1)
	return nil
}

//=====================================================
