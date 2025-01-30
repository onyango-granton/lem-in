package main

import (
	"fmt"

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

	allocatedAnts,_ := functions.AllocateAnts(allPaths, farm.Ants)
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


func SimulateAnts(){
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

	allocatedAnts,paths := functions.AllocateAnts(allPaths, farm.Ants)
	antID := 0
	for i, ch := range allocatedAnts{
		
		for ant:=0; ant<ch; ant++{
			antID++
			fmt.Printf("L%d-path%v\n", antID,paths[i])
		}
	}
}