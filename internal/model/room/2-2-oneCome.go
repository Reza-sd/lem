package room

import "fmt"

// =====================================================
func (u *updater) oneCome() (*room, error) {

	if u.room.get.usedSeats() == u.room.get.allSeats() || u.room.get.usedSeats()+1 > u.room.get.allSeats() {
		return u.room, fmt.Errorf("full")
	}

	u.room.set.usedSeats(u.room.get.usedSeats() + 1)
	return u.room, nil
}

//=====================================================
