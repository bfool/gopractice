package kyu_8

func Reverses(word string) string {
	res := make([]byte, len(word))
	for i, _ := range word {
		res[i] = word[len(word)-i-1]
	}
	return string(res)
}

// interesting
func Reverses2(word string) string {
	var result = ""
	for _, c := range word {
		result = string(c) + result
	}
	return result
}
