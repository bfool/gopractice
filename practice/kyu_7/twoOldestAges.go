package kyu_7

import "sort"

func TwoOldestAges(ages []int) [2]int {
	var res [2]int

	sort.Ints(ages)
	res[0] = ages[len(ages)-2]
	res[1] = ages[len(ages)-1]
	return res
}
