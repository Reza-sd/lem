package room

import "fmt"

// =====================================================
func (myRoom *Room) OneComeUpdateSeats() error {

	if myRoom.UsedSeats == myRoom.AllSeats {
		return fmt.Errorf("full")
	}

	if myRoom.UsedSeats+1 > myRoom.AllSeats {
		return fmt.Errorf("full")
	}

	myRoom.UsedSeats++
	return nil
}

//=====================================================
