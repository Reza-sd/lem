package data

/*
start=0
end=6

0-4
0-6
1-3
4-3
"5-2","3-5","4-2","2-1","7-6","7-2","7-4","6-5"
*/
type Room struct {
	name       string //0 , 1 ,2 ,...
	maxSeats   uint   //all seats in room (for start and end = unlimited) for other =1
	emptySeats uint
	tunnels    []string //[2,5]

}

type Graph struct {
	StartRoomName string
	EndRoomName   string
	Rooms         map[string]Room //string= room name
}
