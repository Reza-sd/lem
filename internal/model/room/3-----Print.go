package room

import "fmt"

func (r *room) Print() {

	fmt.Printf("\nRoom: Name=%v, AllSeats=%v, UsedSeats=%v, ConnectionSlice=%v\n", r.get.name(), r.get.allSeats(), r.get.usedSeats(), r.get.connectionSlice())

}
