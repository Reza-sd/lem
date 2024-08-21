package room

type room struct{
	Name string
	AllSeats Mtr
	FreeSeats Mtr
	TunnelsDb []Mtr
	IsOn bool
}

