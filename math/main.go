package main

import (
	"fmt"
	"time"
)

func main() {
	var x float64 = 1024
	for {
		fmt.Println(x)
		x = x/2
		time.Sleep(2 * time.Second)
	}
}
