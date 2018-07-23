package kyu_6

func Race(v1, v2, g int) (res [3]int) {
	if v1 >= v2 {
		return [3]int{-1, -1, -1}
	}
	s := g * 3600 / (v2 - v1)

	return [3]int{s / 3600, (s / 60) % 60, s % 60}
}
