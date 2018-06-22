package practice

import "strings"

func DuplicateCount(s1 string) (c int) {
	res := make(map[rune]int)

	for _, v := range strings.ToUpper(s1) {
		res[v]++
		if res[v] == 2 {
			c++
		}
	}
	return
}
