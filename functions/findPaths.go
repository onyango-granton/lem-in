package functions

import "fmt"

func BFS(adjList map[string][]string, start, end string) [][]string {
	fmt.Println("HERE", adjList)
	queue  := [][]string{{start}}
	visited := make(map[string]bool)
	visited[start] = true
	var allPaths [][]string

	for len(queue) > 0{
		path := queue[0]
		queue = queue[1:]
		lastRoom := path[len(path)-1]

		if lastRoom == end{
			allPaths = append(allPaths, path)
			continue
		}

		for _, neighbor := range adjList[lastRoom]{
			if !visited[neighbor] {
				
				visited[neighbor] = true
				newPath := append([]string{}, path...)
				
				newPath = append(newPath, neighbor)
				queue = append(queue, newPath)
				
			}
		}
	}
	return allPaths
}