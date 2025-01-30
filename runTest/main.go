package main

import (
	"fmt"
	"lem-in/functions"
)

// func main(){
// 	adjList := map[string][]string{
// 		"A": {"B", "C"},
// 		"B": {"A","D","E"},
// 		"C":{"A","F"},
// 		"D":{"B"},
// 		"E":{"B","F"},
// 		"F":{"C","E"},
// 	}
// 	path := functions.BFS(adjList, "A","F")

// 	paths := [][]string{{"A","B","C"},{"A","D","E","C"},{"A","G","H","F","C"}}

// 	fmt.Println(functions.AllocateAnts(paths, 6))

// 	fmt.Println(path)
// }

func main() {
	farm, err := functions.ParseInput("../example.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	path := functions.BFS(farm.AdjList, farm.Start, farm.End)
	if path == nil{
		fmt.Println("No path found")
		return
	}

	antsPerPath := functions.AllocateAnts([][]string{path}, farm.Ants)
	functions.SimulateAnts([][]string{path}, antsPerPath)
}