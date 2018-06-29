package practice

import "strings"
import "strconv"

func IsValidCoordinates(coordinates string) (res bool) {
	var arr []string = strings.Split(coordinates, ", ")

	if len(arr) > 2 || strings.Contains(coordinates, "e") {
		return false
	}
	a, err_a := strconv.ParseFloat(arr[0], 64)
	b, err_b := strconv.ParseFloat(arr[1], 64)
	if err_a != nil || err_b != nil {
		return false
	} else {
		res = judge(a, b)
	}

	return res
}

func judge(a, b float64) bool {
	if ((a > 0.0 && a < 90.0) || (a > -90.0 && a < -0.0)) && ((b > 0.0 && b < 180.0) || (b > -180.0 && b < -0.0)) {
		return true
	}
	return false
}
