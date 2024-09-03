package room

import "fmt"

// =====================================================
func (r *room) OneComeUpdateSeats() error {

	if r.get.usedSeats() == r.get.allSeats() {
		return fmt.Errorf("full")
	}

	if r.get.usedSeats()+1 > r.get.allSeats() {
		return fmt.Errorf("full")
	}

	//r.usedSeats++
	r.set.usedSeats(r.get.usedSeats() + 1)
	return nil
}

//=====================================================
