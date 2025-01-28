package main

import (
	"fmt"
	"os"
)

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

func readFile(fileName string) string{
	contents, err := os.ReadFile(fileName)
	if err != nil {
		return err.Error()
	}
	return string(contents)
}

func main(){
	fmt.Println(readFile("example.txt"))
}