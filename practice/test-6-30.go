package practice

import "strconv"

func Test_6_30(str string) string {
	res := ""
	temp := 1
	for i, _ := range str {
		if res == "" {
			res += string(str[i])
		} else if str[i] == res[len(res)-1] {
			temp += 1
		} else {
			res += strconv.Itoa(temp)
			res += string(str[i])
			temp = 1
		}
	}
	res += strconv.Itoa(temp)
	return res
}
