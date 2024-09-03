package room

import "fmt"

// =====================================================

func (myRoom *Room) OneLeaveUpdateSeats() error {

	if myRoom.usedSeats == 0 {
		return fmt.Errorf("full")
	}
	myRoom.usedSeats--
	return nil
}

// =====================================================
