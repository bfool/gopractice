package kyu_7

// import "math"

func EquableTriangle(a, b, c int) bool {
	// p := float64(a+b+c) / 2.0
	// area := math.Sqrt(p * (p - float64(a)) * (p - float64(b)) * (p - float64(c)))
	// return float64(a+b+c) == area

	p := a + b + c
	return 16*p == (p-2*a)*(p-2*b)*(p-2*c)
}
