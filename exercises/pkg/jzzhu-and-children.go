package pkg

import (
	"math"
	"slices"
)

type Queue struct {
	value []Child
}

type Child struct {
	place int
	rem   int
}

func (q *Queue) Enqueue(item Child) {
	q.value = append(q.value, item)
}

func (q *Queue) Unqueue() Child {
	item := q.value[0]
	q.value = q.value[1:]
	return item
}

func JzzhuAndChildrenV1(n, m int, ai []int) int {
	q := Queue{}
	for i, v := range ai {
		q.Enqueue(Child{
			place: i + 1,
			rem:   v,
		})
	}

	var ch Child
	for len(q.value) != 0 {
		ch = q.Unqueue()
		ch.rem = ch.rem - m
		if ch.rem > 0 {
			q.Enqueue(ch)
		}
	}
	return ch.place
}

func JzzhuAndChildrenV2(n, m int, ai []int) int {
	q := Queue{}
	for i, v := range ai {
		if v > m {
			q.Enqueue(Child{
				place: i + 1,
				rem:   v,
			})
		}
	}

	if len(q.value) == 0 {
		return n
	}

	var ch Child
	for len(q.value) != 0 {
		ch = q.Unqueue()
		ch.rem = ch.rem - m
		if ch.rem > 0 {
			q.Enqueue(ch)
		}
	}
	return ch.place
}

func JzzhuAndChildrenV3(n, m int, ai []int) int {
	s := make([]struct {
		index int
		val   int
	}, n)
	var mult int
	for i, v := range ai {
		mult = v / m
		ai[i] = mult
		s[i] = struct {
			index int
			val   int
		}{index: i + 1, val: mult}
	}
	maxInt := slices.Max(ai)
	slices.Reverse(s)
	indexFounded := slices.IndexFunc(s, func(e struct {
		index int
		val   int
	}) bool {
		return e.val == maxInt
	})
	return s[indexFounded].index
}

func JzzhuAndChildrenV4(n, m int, ai []int) int {
	loop := -1
	index := 0
	var c int
	for i := 0; i < n; i++ {
		c = int(math.Ceil(float64(ai[i]) / float64(m)))
		if loop <= c {
			loop = c
			index = i + 1
		}
	}
	return index
}
