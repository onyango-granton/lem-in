package main

//ant struct
//room struct

type ant struct{
	ant_number int
	roomsVisited *[]room
}

type room struct{
	name string
	coord_x int
	coord_y int
	roomLink *link
	visited bool
}

type link struct{
	room_before *room
	room_after *room
}

func main(){
	
}