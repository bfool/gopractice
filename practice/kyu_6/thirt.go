package kyu_6

import "strconv"

var nums = [6]int{1, 10, 9, 12, 3, 4}

func Thirt(n int) int {
	preview := get_sum(n)
	for {
		next := get_sum(preview)
		if preview == next {
			return preview
		} else {
			preview = next
		}
	}
	return 0
}

func get_sum(n int) int {
	var sum int
	j := 0
	s := strconv.Itoa(n)
	for i := len(s) - 1; i > -1; i-- {
		num, _ := strconv.Atoi(string(s[i]))
		sum += num * nums[j%6]
		j++
	}
	return sum
}
