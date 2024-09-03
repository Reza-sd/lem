package room

import "fmt"

func (r *Room) Print() {

	fmt.Printf("\nRoom: Name=%v, AllSeats=%v, UsedSeats=%v, ConnectionSlice=%v\n", r.getName(), r.getAllSeats(), r.getUsedSeats(), r.getConnectionSlice())

}
