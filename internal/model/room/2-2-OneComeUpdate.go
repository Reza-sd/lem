package room

import "fmt"


func (myRoom *Room)OneComeUpdate () error {

	if myRoom.UsedSeats+1>myRoom.AllSeats {return fmt.Errorf("full")}
	myRoom.UsedSeats++
	return nil
}