package practice

import "sort"

func convertTo64(arr []float32) []float64 {
	newer := make([]float64, len(arr))
	for i, v := range arr {
		newer[i] = float64(v)
	}
	return newer
}

func FindUniq(arry []float32) float32 {
	arr := convertTo64(arry)

	sort.Float64s(arr)
	if arr[0] < arr[len(arr)-1] && arr[0] < arr[len(arr)-2] {
		return float32(arr[0])
	} else {
		return float32(arr[len(arr)-1])
	}

}
