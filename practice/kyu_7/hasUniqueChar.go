package kyu_7

func HasUniqueChar(str string) bool {
	temp := make(map[rune]int)
	for _, v := range str {
		if temp[v]++; temp[v] > 1 {
			return false
		}
	}
	return true
}
