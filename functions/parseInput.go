package functions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseInput(filename string) (Farm, error){
	file, err := os.Open(filename)
	if err != nil {
		return Farm{}, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	farm := Farm{AdjList: make(map[string][]string)}
	var isStart, isEnd bool

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "##") {
			if line == "##start"{
				isStart = true
			} else if line == "##end"{
				isEnd = true
			}
			continue
		}

		if farm.Ants == 0{
			//obtaining ants number given they should be parsed in line 1 of file
			ants, err := strconv.Atoi(line)
			if err != nil {
				return Farm{}, fmt.Errorf("invalid ant count: %v",err)
			}
			farm.Ants = ants
		}else if strings.Contains(line, " "){
			// obtaining rooms given thay are uniquely identified with spaces
			parts := strings.Split(line, " ")
			room := parts[0]
			farm.Rooms = append(farm.Rooms, room)
			if isStart {
				farm.Start = room
				isStart = false
			} else if isEnd{
				farm.End = room
				isEnd = false
			}
		} else if strings.Contains(line, "-"){
			//obtaining links given thay are uniquely identified with "hyphen"
			farm.Links = append(farm.Links, line)
			parts := strings.Split(line, "-")
			farm.AdjList[parts[0]] = append(farm.AdjList[parts[0]], parts[1])
			farm.AdjList[parts[1]] = append(farm.AdjList[parts[1]], parts[0])

		}
	}

	return farm, nil
}