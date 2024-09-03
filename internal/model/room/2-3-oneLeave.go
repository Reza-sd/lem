package room

import "fmt"

// =====================================================
func (u *updater) oneLeave() (*room, error) {
	if u.room.get.usedSeats() == 0 {
		return u.room, fmt.Errorf("full")
	}
	u.room.set.usedSeats(u.room.get.usedSeats() - 1)
	return u.room, nil
}

// =====================================================
