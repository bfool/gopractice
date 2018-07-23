package kyu_6

var num map[rune]int = map[rune]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

func Decode(roman string) (res int) {
	for i, v := range roman {
		if i+1 < len(roman) && num[v] < num[rune(roman[i+1])] {
			res -= num[v]
		} else {
			res += num[v]
		}
	}
	return
}
