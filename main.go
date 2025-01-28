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

func buildGraph(routes []string) map[string][]string{
	graph := make(map[string][]string)
	for _, route := range routes{
		route := strings.Split(route, "-")
		src, dest := route[0], route[1]
		graph[src] = append(graph[src], dest)
	}
	return graph
}

func findRoutes(graph map[string][]string,current, end string, pathRoute []string, allPath *[][]string){
	pathRoute = append(pathRoute, current)

	if current == end{
		*allPath = append(*allPath, append([]string{}, pathRoute...))
		return
	}

	for _, neighbor := range graph[current]{
		findRoutes(graph, neighbor, end, pathRoute, allPath)
	}
}

func main(){
	// fmt.Println(getLinks("example.txt"))
	fmt.Println(buildGraph(getLinks("example.txt")))
}