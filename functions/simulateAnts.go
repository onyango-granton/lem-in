package functions

import (
	"fmt"
	"strings"
)

func SimulateAnts(paths [][]string, allocatedAnts []int) {
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
