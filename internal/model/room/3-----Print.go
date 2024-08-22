package room

import "fmt"

func (r *Room) Print() {

	fmt.Printf("\nRoom: Name=%v, AllSeats=%v, UsedSeats=%v, ConnectionSlice=%v\n", r.Name, r.AllSeats, r.UsedSeats, r.ConnectionSlice)

}
