package main

import (
	"fmt"
)

// Edmonds-Karp algorithm to find maximum flow in a graph
func edmondsKarp(graph [][]int, source, sink int) int {
	n := len(graph)
	parent := make([]int, n)
	maxFlow := 0

	// BFS to find augmenting paths
	bfs := func() bool {
		visited := make([]bool, n)
		queue := []int{source}
		visited[source] = true

		for len(queue) > 0 {
			u := queue[0]
			queue = queue[1:]

			for v := 0; v < n; v++ {
				if !visited[v] && graph[u][v] > 0 {
					queue = append(queue, v)
					parent[v] = u
					visited[v] = true
				}
			}
		}

		return visited[sink]
	}

	// Find the maximum flow
	for bfs() {
		pathFlow := 1 << 30 // Initialize path flow with a large value
		v := sink

		// Find the minimum residual capacity of the path
		for v != source {
			u := parent[v]
			if graph[u][v] < pathFlow {
				pathFlow = graph[u][v]
			}
			v = u
		}

		// Update residual capacities and reverse edges
		v = sink
		for v != source {
			u := parent[v]
			graph[u][v] -= pathFlow
			graph[v][u] += pathFlow
			v = u
		}

		// Add path flow to the total flow
		maxFlow += pathFlow
	}

	return maxFlow
}

func main() {
	// Example graph represented as an adjacency matrix
	graph := [][]int{
		{0, 16, 13, 0, 0, 0},
		{0, 0, 10, 12, 0, 0},
		{0, 4, 0, 0, 14, 0},
		{0, 0, 9, 0, 0, 20},
		{0, 0, 0, 7, 0, 4},
		{0, 0, 0, 0, 0, 0},
	}

	source := 0
	sink := 5

	maxFlow := edmondsKarp(graph, source, sink)
	fmt.Println("Maximum flow:", maxFlow) // Expected output: 23
}