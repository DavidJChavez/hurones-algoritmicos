package main

import (
	"algorithms/pkg"
	"algorithms/utils"
	"fmt"
	"math/rand"
	"testing"
)

var table = []struct {
	name     string
	n, m     int
	ai       []int
	expected int
}{
	{name: "Case 1", n: 5, m: 2, ai: []int{1, 3, 1, 4, 2}, expected: 4},
	{name: "Case 2", n: 6, m: 4, ai: []int{1, 1, 2, 2, 3, 3}, expected: 6},
}

func TestJzzhuAndChildren(t *testing.T) {
	for _, v := range table {
		t.Run(v.name, func(t *testing.T) {
			r := pkg.JzzhuAndChildrenV4(v.n, v.m, v.ai)
			if r == v.expected {
				utils.PrintSuccess("Passed!")
			} else {
				utils.PrintErr(fmt.Sprintf("Failed! Expected %d but received %d ", v.expected, r))
			}
		})
	}
}

func BenchmarkJzzhuAndChildren(b *testing.B) {
	var n, m int
	var ai []int
	minN := 1
	minM := 1
	maxN := 100
	maxM := 100
	minA := 1
	maxA := 100
	b.Run("Testing cases", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			n = rand.Intn(maxN-minN+1) + minN
			m = rand.Intn(maxM-minM+1) + minM
			ai = make([]int, n)
			for i := 0; i < n; i++ {
				ai[i] = rand.Intn(maxA-minA+1) + minA
			}
			pkg.JzzhuAndChildrenV4(n, m, ai)
		}
	})
}
