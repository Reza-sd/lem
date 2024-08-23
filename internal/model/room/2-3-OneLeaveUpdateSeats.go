package room

import "fmt"

// =====================================================

func (myRoom *Room) OneLeaveUpdateSeats() error {

	if myRoom.UsedSeats == 0 {
		return fmt.Errorf("full")
	}
	myRoom.UsedSeats--
	return nil
}

// =====================================================
