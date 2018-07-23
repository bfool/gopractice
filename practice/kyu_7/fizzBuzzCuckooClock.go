package kyu_7

import (
	"strconv"
	"strings"
)

func FizzBuzzCuckooClock(time string) string {
	var arr []string = strings.Split(time, ":")
	hour, _ := strconv.Atoi(arr[0])
	minutes, _ := strconv.Atoi(arr[1])

	// switch {
	// case minutes == 0:
	// 	return strings.Repeat("Cuckoo ", (hour+11)%12) + "Cuckoo"
	// case m == 30:
	// 	return "Cuckoo"
	// case m%15 == 0:
	// 	return "Fizz Buzz"
	// case m%5 == 0:
	// 	return "Buzz"
	// case m%3 == 0:
	// 	return "Fizz"
	// default:
	// 	return "tick"
	// }
	if minutes == 0 {
		times := (hour + 11) % 12
		var res string
		for i := 0; i < times; i++ {
			res += "Cuckoo "
		}
		return res + "Cuckoo"
	}
	if minutes == 30 {
		return "Cuckoo"
	}
	if minutes%3 == 0 && minutes%5 == 0 {
		return "Fizz Buzz"
	}
	if minutes%3 == 0 {
		return "Fizz"
	}
	if minutes%5 == 0 {
		return "Buzz"
	}
	return "tick"
}
