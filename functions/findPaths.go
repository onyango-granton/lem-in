package functions

func BFS(adjList map[string][]string, start, end string) []string {
	queue  := [][]string{{start}}
	visited := make(map[string]bool)
	visited[start] = true

	for len(queue) > 0{
		path := queue[0]
		queue = queue[1:]
		lastRoom := path[len(path)-1]

		if lastRoom == end{
			return path
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
	return nil
}