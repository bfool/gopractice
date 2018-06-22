package practice

import "unicode"

func Accum(s string) string {
	var newer []byte
	for i, v := range s {
		for j := 0; j < i+1; j++ {
			if j == 0 {
				newer = append(newer, byte(unicode.ToUpper(v)))
			} else {
				newer = append(newer, byte(unicode.ToLower(v)))
			}
		}
		newer = append(newer, '-')
	}

	return string(newer[:len(newer)-1])
}
