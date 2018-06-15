package practice

// remeber
// when it can solve by fool style, just do it
// and optimization it, don't do nothing
func SumEvenFibonacci(limit int) int {
	a, b, sum := 1, 2, 2

	for b <= limit {
		a, b = b, a+b
		if b%2 == 0 {
			sum += b
		}
	}
	return sum
}
