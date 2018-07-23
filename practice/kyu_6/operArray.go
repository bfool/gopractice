package kyu_6

func Gcdi(x, y int) int {
	// tmp := x % y
	// if tmp > 0 {
	// 	return Gcdi(y, tmp)
	// } else {
	// 	return y
	// }
	var tmp int
	x = abs(x)
	y = abs(y)
	for {
		tmp = (x % y)
		if tmp > 0 {
			x = y
			y = tmp
		} else {
			return y
		}
	}
}
func Som(x, y int) int {
	return x + y
}
func Maxi(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func Mini(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func Lcmu(x, y int) int {
	return (abs(x) * abs(y)) / Gcdi(abs(x), abs(y))
}

func abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}

type FParam func(int, int) int

func OperArray(f FParam, arr []int, init int) (res []int) {
	res = append(res, f(arr[0], init))
	for i := 1; i < len(arr); i++ {
		res = append(res, f(res[i-1], arr[i]))
	}
	return res
}
