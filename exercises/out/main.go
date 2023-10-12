package main

import (
	"fmt"
	"math"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	loop := -1
	index := 0
	var c, child int
	for i := 0; i < n; i++ {
		fmt.Scan(&child)
		c = int(math.Ceil(float64(child) / float64(m)))
		if loop <= c {
			loop = c
			index = i + 1
		}
	}
	fmt.Println(index)
}
