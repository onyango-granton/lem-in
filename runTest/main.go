package main

import (
	"fmt"
	"strings"

	"lem-in/functions"
)

func main() {
	farm, err := functions.ParseInput("example.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Call DFS to find all paths
	allPaths := functions.DFS(farm.AdjList, farm.Start, farm.End)
	if len(allPaths) == 0 {
		fmt.Println("No paths found")
		return
	}

	// Print all paths
	for _, path := range allPaths {
		fmt.Println(path)
	}

	allocatedAnts, _ := functions.AllocateAnts(allPaths, farm.Ants)
	functions.SimulateAnts(allPaths, allocatedAnts)
	SimulateAnts()
}

func SimulateAnts() {
	farm, err := functions.ParseInput("example.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Call DFS to find all paths
	allPaths := functions.DFS(farm.AdjList, farm.Start, farm.End)
	if len(allPaths) == 0 {
		fmt.Println("No paths found")
		return
	}

	// Print all paths
	for _, path := range allPaths {
		fmt.Println(path)
	}

	// get longest path
	// for i, ch := range allocatedAnts{

	// 	for ant:=0; ant<ch; ant++{
	// 		antID++
	// 		antPath[antID]=paths[i]
	// 	}
	// }

	// for _,path := range paths{
	// 	//fmt.Println(index, path)
	// 	for i, step := range path{
	// 		for j, _ := range antPath{
	// 			if i >= len(antPath[j]){
	// 				break
	// 			}
	// 			if step == antPath[j][i]{
	// 				fmt.Printf("L%v%v",j,step)
	// 				break
	// 			}
	// 		}

	// 	}
	// }

	findPaths()
}

func findPaths() {
	farm, err := functions.ParseInput("example2.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Call DFS to find all paths
	allPaths := functions.DFS(farm.AdjList, farm.Start, farm.End)
	if len(allPaths) == 0 {
		fmt.Println("No paths found")
		return
	}

	allocatedAnts, paths := functions.AllocateAnts(allPaths, farm.Ants)
	// paths := [][]string{
	// 	{"start", "t", "E", "a", "m", "end"},
	// 	{"start", "h", "A", "c", "k", "end"},
	// 	{"start", "h", "n", "e", "end"},
	// 	{"start", "h", "n", "m", "end"},
	// 	{"start", "0", "o", "n", "e", "end"},
	// 	{"start", "0", "o", "n", "m", "end"},
	// }
	// allocatedAnts := []int{2, 2, 2, 2, 1, 1}

	// Assign unique IDs to each ant based on path and order
	currentAntID := 1
	antIDs := make([][]int, len(paths))
	for pathIdx := range paths {
		antCount := allocatedAnts[pathIdx]
		antIDs[pathIdx] = make([]int, antCount)
		for i := 0; i < antCount; i++ {
			antIDs[pathIdx][i] = currentAntID
			currentAntID++
		}
	}

	// Calculate maximum time required for all ants to finish
	maxTime := 0
	for pathIdx := range paths {
		pathLen := len(paths[pathIdx]) - 1 // exclude 'start'
		antCount := allocatedAnts[pathIdx]
		for antInPath := 0; antInPath < antCount; antInPath++ {
			if t := antInPath + pathLen; t > maxTime {
				maxTime = t
			}
		}
	}

	// Simulate each time step and collect movements
	for time := 1; time <= maxTime; time++ {
		var moves []string
		for pathIdx := range paths {
			antCount := allocatedAnts[pathIdx]
			path := paths[pathIdx]
			for antInPath := 0; antInPath < antCount; antInPath++ {
				step := time - antInPath
				if step < 1 || step >= len(path) {
					continue // Ant hasn't reached or passed this step
				}
				antID := antIDs[pathIdx][antInPath]
				moves = append(moves, fmt.Sprintf("L%d-%s", antID, path[step]))
			}
		}
		if len(moves) > 0 {
			fmt.Println(strings.Join(moves, " "))
		}
	}
}
