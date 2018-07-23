package kyu_7

func CartesianNeighbor(x, y int) (res [][]int) {
	// var coords [][]int
	// for i := -1; i <= 1; i++ {
	//   for j := -1; j <= 1; j++ {
	//     dx := x + i
	//     dy := y + j
	//     if !(dx == x && dy == y) {
	//       coords = append(coords, []int{dx, dy})
	//     }
	//   }
	// }
	// return coords
	return [][]int{[]int{x - 1, y + 1}, []int{x, y + 1}, []int{x + 1, y + 1},
		[]int{x - 1, y}, []int{x + 1, y},
		[]int{x - 1, y - 1}, []int{x, y - 1}, []int{x + 1, y - 1}}

}
