package room

import "fmt"

// =====================================================
func (r *Room) OneLeaveUpdateSeats() error {
	if r.getUsedSeats() == 0 {
		return fmt.Errorf("full")
	}
	r.setUsedSeats(r.getUsedSeats() - 1)
	return nil
}

// =====================================================
