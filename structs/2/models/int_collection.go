package models

import "sort"

type Collection []int

func (c Collection) Sort() Collection {
	sorted := make([]int, len(c))
	copy(sorted, c)
	sort.Ints(sorted)
	return sorted
}

func (c Collection) Max() int {
	max := c[0]
	for _, item := range c {
		if item > max {
			max = item
		}
	}

	return max
}

func (c Collection) Min() int {
	min := c[0]
	for _, item := range c {
		if item < min {
			min = item
		}
	}

	return min
}

func (c Collection) Evens() Collection {
	col := make(Collection, 0)

	for _, item := range c {
		if item % 2 == 0 {
			col = append(col, item)
		}
	}

	return col
}
