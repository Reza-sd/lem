package room

import "fmt"

// =====================================================
func (r *Room) OneLeaveUpdateSeats() error {
	if r.get.usedSeats() == 0 {
		return fmt.Errorf("full")
	}
	r.set.usedSeats(r.get.usedSeats() - 1)
	return nil
}

// =====================================================
