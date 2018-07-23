package kyu_8

func MultipleOfIndex(ints []int) (res []int) {
	for i, v := range ints {
		if i == 0 {
			continue
		} else if v%i == 0 {
			res = append(res, v)
		}
	}
	return res
}
