package kyu_8

func Century(year int) int {
	// res := float64(year) / 100.0
	// if res-float64(int(res)) > 0 {
	// 	res += 1
	// }
	// return int(res)

	// woooooooooooooooo!
	return (year + 99) / 100
}
