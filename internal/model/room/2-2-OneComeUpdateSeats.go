package room

import "fmt"

// =====================================================
func (r *Room) OneComeUpdateSeats() error {

	if r.getUsedSeats() == r.getAllSeats() {
		return fmt.Errorf("full")
	}

	if r.getUsedSeats()+1 > r.getAllSeats() {
		return fmt.Errorf("full")
	}

	//r.usedSeats++
	r.setUsedSeats(r.getUsedSeats() + 1)
	return nil
}

//=====================================================
