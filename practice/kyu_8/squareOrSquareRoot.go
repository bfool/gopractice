package kyu_8

import "math"

func SquareOrSquareRoot(arr []int) (res []int) {
	for _, v := range arr {
		t := math.Sqrt(float64(v))
		if t == float64(int(t)) {
			res = append(res, int(t))
		} else {
			res = append(res, v*v)
		}
	}
	return res
}
