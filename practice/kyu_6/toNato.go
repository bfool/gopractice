package kyu_6

import "strings"

func ToNato(words string) string {
	nato := map[string]string{"A": "Alfa", "B": "Bravo", "C": "Charlie", "D": "Delta",
		"E": "Echo", "F": "Foxtrot", "G": "Golf", "H": "Hotel",
		"I": "India", "J": "Juliett", "K": "Kilo", "L": "Lima",
		"M": "Mike", "N": "November", "O": "Oscar", "P": "Papa",
		"Q": "Quebec", "R": "Romeo", "S": "Sierra", "T": "Tango",
		"U": "Uniform", "V": "Victor", "W": "Whiskey", "X": "X-ray",
		"Y": "Yankee", "Z": "Zulu"}

	var res string
	for _, v := range strings.ToUpper(words) {
		if string(v) == " " {
			continue
		}
		if _, ok := nato[string(v)]; ok {
			res += nato[string(v)] + " "
		} else {
			res += string(v)
		}
	}
	if string(res[len(res)-1]) == " " {
		res = res[:len(res)-1]
	}

	return res
}
