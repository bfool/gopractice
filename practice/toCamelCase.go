package practice

import (
	"regexp"
	"unicode"
)

func ToCamelCase(s string) string {
	newer := make([]byte, len(s))
	preview := false

	for i, _ := range s {
		if preview == true {
			preview = false
			continue
		}
		is_ok, _ := regexp.MatchString(`[A-Za-z]+`, string(s[i]))
		if is_ok == false && i == len(s)-1 {
			break
		}
		if is_ok == false {
			preview = true
			newer[i] = byte(unicode.ToUpper(rune(s[i+1])))
		} else {
			newer[i] = s[i]
		}
	}
	return string(newer)
}
