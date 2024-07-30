package graphpk

/*
"0-4","0-6","1-3","4-3","5-2","3-5","4-2","2-1","7-6","7-2","7-4","6-5"

L1-3 L2-2
L1-4 L2-5 L3-3
L1-0 L2-6 L3-4
L2-0 L3-0
$
*/
type Room struct {
	name       string //0 , 1 ,2 ,...
	maxSeats   uint   //all seats in room (for start and end = unlimited) for other =1
	EmptySeats uint
	Tunnels    []string //[2,5]

}

type Graph struct {
	StartRoomName        string
	EndRoomName          string
	CurrentAntsInEndRoom int
	Rooms                map[string]Room //string= room name
	//Paths                map[int][][]string
}
