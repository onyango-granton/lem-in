package main

import (
	"fmt"
	"os"
	"strings"
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

func getLinks(fileName string) []string{
	var links []string
	rooms := strings.Split(readFile(fileName), string(rune(10)))
	for _,ch := range rooms{
		if strings.Contains(ch, "-"){
			links = append(links, ch)
		}
	}

	return links
}

func main(){
	fmt.Println(getLinks("example.txt"))
}