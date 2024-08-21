package room

import "fmt"

func (myRoom *Room) OneLeaveUpdate() error {

	if myRoom.UsedSeats == 0 {
		return fmt.Errorf("full")
	}
	myRoom.UsedSeats--
	return nil
}
