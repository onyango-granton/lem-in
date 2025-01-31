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
	allPaths := DFS(farm.AdjList, farm.Start, farm.End)
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

// DFS function to find all paths from start to end
func DFS(adjList map[string][]string, start, end string) [][]string {
	visited := make(map[string]bool) // Track visited rooms to avoid cycles
	var allPaths [][]string          // Store all valid paths
	var currentPath []string         // Track the current path being explored

	// Recursive DFS helper function
	var dfsHelper func(node string)
	dfsHelper = func(node string) {
		// Mark the current node as visited and add it to the current path
		visited[node] = true
		currentPath = append(currentPath, node)

		// If the current node is the end, add the current path to allPaths
		if node == end {
			// Make a copy of the current path to avoid overwriting
			pathCopy := make([]string, len(currentPath))
			copy(pathCopy, currentPath)
			allPaths = append(allPaths, pathCopy)
		} else {
			// Explore all neighbors of the current node
			for _, neighbor := range adjList[node] {
				if !visited[neighbor] {
					dfsHelper(neighbor)
				}
			}
		}

		// Backtrack: Remove the current node from the path and mark it as unvisited
		currentPath = currentPath[:len(currentPath)-1]
		visited[node] = false
	}

	// Start DFS from the start node
	dfsHelper(start)

	return allPaths
}

func SimulateAnts() {
	farm, err := functions.ParseInput("example.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Call DFS to find all paths
	allPaths := DFS(farm.AdjList, farm.Start, farm.End)
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
	allPaths := DFS(farm.AdjList, farm.Start, farm.End)
	if len(allPaths) == 0 {
		fmt.Println("No paths found")
		return
	}

	allocatedAnts, paths :=functions.AllocateAnts(allPaths, farm.Ants)
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
