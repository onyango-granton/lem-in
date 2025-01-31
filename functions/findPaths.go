package functions

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
