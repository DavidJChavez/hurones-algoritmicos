package main

import (
	"algorithms/pkg"
	"fmt"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	ai := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&ai[i])
	}
	fmt.Println(pkg.JzzhuAndChildrenV1(n, m, ai))
}
