package room

import "fmt"

// =====================================================
func (myRoom *Room) OneComeUpdateSeats() error {

	if myRoom.usedSeats == myRoom.allSeats {
		return fmt.Errorf("full")
	}

	if myRoom.usedSeats+1 > myRoom.allSeats {
		return fmt.Errorf("full")
	}

	myRoom.usedSeats++
	return nil
}

//=====================================================
