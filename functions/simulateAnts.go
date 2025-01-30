package functions

import "fmt"

func SimulateAnts(paths [][]string, antsPerPath []int) {
	antPositions := make([]int, len(paths))
	for step := 1; ; step++ {
		moved := false
		for i := 0; i < len(paths); i++ {
			if antPositions[i] < len(paths[i]) && antsPerPath[i] > 0 {
				fmt.Printf("L%d-%s ", step, paths[i][antPositions[i]])
				antPositions[i]++
				antsPerPath[i]--
				moved = true
			}
		}
		if !moved {
			break
		}
		fmt.Println()
	}
}