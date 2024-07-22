package main

type room struct{
	name uint //0 , 1 ,2 ,...
	maxSeats uint //all seats in room (for start and end = unlimited) for other =1 
	emptySeats uint
	tunnels []uint //[2,5]

}

type graph struct{
	startRoomName uint
	endRoomName uint
	rooms map[uint]room //uint= room name
}