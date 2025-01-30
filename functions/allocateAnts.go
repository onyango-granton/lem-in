package functions

import "fmt"

func AllocateAnts(paths [][]string, ants int) ([]int,[][]string) {
	antsPerPath := make([]int, len(paths))
	pathLengths := make([]int, len(paths))

	for i, path := range paths {
		pathLengths[i] = len(path)
	}

	for i := 0; i < ants; i++{
		minIndex := 0
		minValue := antsPerPath[0]+pathLengths[0]

		for j:=0; j < len(paths); j++{
			if antsPerPath[j]+pathLengths[j] < minValue{
				minIndex = j
				minValue = antsPerPath[j] + pathLengths[j]
			}
		}
		antsPerPath[minIndex]++
	}
	fmt.Println(antsPerPath, paths)
	return antsPerPath, paths
}