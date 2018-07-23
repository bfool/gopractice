package kyu_7

import "strings"
import "regexp"

func BandNameGenerator(word string) string {
	str := make([]string, len(word))
	temp := word[0]
	str[0] = string(word[0])
	for i := 1; i < len(word); i++ {
		is_ok, _ := regexp.MatchString(`[A-Za-z]+`, string(temp))
		if is_ok {
			str[i] = string(word[i])
		} else {
			str[i] = strings.ToUpper(string(word[i]))
		}
		temp = word[i]
	}
	word = strings.Join(str, "")
	if word[0] == word[len(word)-1] {
		return strings.ToUpper(string(word[0])) + word[1:len(word)] + word[1:len(word)]
	} else {
		return "The " + strings.ToUpper(string(word[0])) + word[1:len(word)]
	}
}

func Easy(word string) string {
	first := word[:1]
	last := word[len(word)-1:]

	if first != last {
		return "The " + strings.Title(word)
	}

	return strings.Title(word) + word[1:]
}
