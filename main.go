package main

import (
	"fmt"

	"lem-in/functions"
)

func main() {
	farm, err := functions.ParseInput("example2.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Call DFS to find all paths
	allPaths := functions.FindPathsDFS(farm.AdjList, farm.Start, farm.End)
	if len(allPaths) == 0 {
		fmt.Println("No paths found")
		return
	}

	allocatedAnts := functions.AllocateAnts(allPaths, farm.Ants)
	functions.SimulateAnts(allPaths, allocatedAnts)
}
