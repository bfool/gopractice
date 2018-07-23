package kyu_7

// import "sort"

func IsTriangle(a, b, c int) bool {
	// if a < 0 || b < 0 || c < 0 {
	// 	return false
	// }
	//
	// arr := []int{a, b, c}
	// sort.Ints(arr)
	// return arr[0]+arr[1] > arr[2]
	//
	return a+b > c && b+c > a && a+c > b
}
