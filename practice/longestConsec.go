package practice

func LongestConsec(strarr []string, k int) (res string) {
	length := len(strarr)
	if length == 0 || k > length || k <= 0 {
		return ""
	}

	for i := 0; i < len(strarr)-1; i++ {
		for j := i + 1; j < len(strarr); j++ {
			if len(strarr[i]) < len(strarr[j]) {
				strarr[i], strarr[j] = strarr[j], strarr[i]
			}
		}
	}

	for i := 0; i < k; i++ {
		res += strarr[i]
	}
	return
}
