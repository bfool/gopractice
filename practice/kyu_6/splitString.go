package kyu_6

func SplitString(str string) []string {
	var res []string
	for i, v := range str {
		if i%2 == 0 {
			res = append(res, string(v))
		} else {
			res[len(res)-1] += string(v)
		}
	}
	if len(res[len(res)-1]) == 1 {
		res[len(res)-1] += "_"
	}
	return res
}

func Solution2(str string) (res []string) {
	if len(str)%2 != 0 {
		str += "_"
	}

	for i := 0; i < len(str); i += 2 {
		res = append(res, str[i:i+2])
	}

	return
}
