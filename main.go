package main

import (
	"fmt"
	"gopractice/practice"
	"gopractice/practice/kyu_5"
	"gopractice/practice/kyu_6"
	"gopractice/practice/kyu_7"
	"gopractice/practice/kyu_8"
)

func validParentheses() {
	// var testString string
	// fmt.Scanln(&testString)
	// fmt.Println("InputSring: ", testString)

	res := kyu_5.StackValidParentheses("{}[]<>(")
	fmt.Println(res)
}

func sumEvenFibonacci() {
	res := kyu_7.SumEvenFibonacci(1)
	fmt.Println(res)
}

func findUniq() {
	s := []float32{1, 2, 1}
	res := kyu_6.FindUniq(s)
	fmt.Println(res)
}

func toCamelCase() {
	s := "th"
	res := kyu_6.ToCamelCase(s)
	fmt.Println(res)
}

func twoOldestAges() {
	s := []int{6, 5, 83, 5, 3, 18}
	res := kyu_7.TwoOldestAges(s)
	fmt.Println(res)
}

func multiple3And5() {
	res := kyu_6.Multiple3And5(10)
	fmt.Println(res)
}

func evenOrOdd() {
	res := kyu_8.EvenOrOdd(2)
	fmt.Println(res)
}

func positiveSum() {
	res := kyu_8.PositiveSum([]int{1, -2, 3, 4, 5})
	fmt.Println(res)
}

func removeChar() {
	res := kyu_8.RemoveChar("person")
	fmt.Println(res)
}

func repeatStr() {
	res := kyu_8.RepeatStr(4, "hello ")
	fmt.Println(res)
}

func century() {
	res := kyu_8.Century(89)
	fmt.Println(res)
}

func isTriangle() {
	res := kyu_7.IsTriangle(2, 5, 1)
	fmt.Println(res)
}

func accum() {
	res := kyu_7.Accum("RqaEzty")
	fmt.Println(res)
}
func duplicateCount() {
	res := kyu_6.DuplicateCount("inddivisibility")
	fmt.Println(res)
}

func inAscOrder() {
	res := kyu_7.InAscOrder([]int{1, 2, 4, 7, 19})
	fmt.Println(res)
}

func reverses() {
	res := kyu_8.Reverses("hello")
	fmt.Println(res)
}

func longestConsec() {
	res := kyu_7.LongestConsec([]string{"itvayloxrp", "wkppqsztdkmvcuwvereiupccauycnjutlv", "vweqilsfytihvrzlaodfixoyxvyuyvgpck"}, 2)
	fmt.Println(res)
}

func toNato() {
	res := kyu_6.ToNato("go for it!")
	fmt.Println(res)
}

func decode() {
	res := kyu_6.Decode("MCMXC")
	fmt.Println(res)
}

func hasUniqueChar() {
	res := kyu_7.HasUniqueChar("+-")
	fmt.Println(res)
}

func bandNameGenerator() {
	res := kyu_7.BandNameGenerator("tar-zc")
	fmt.Println(res)
}

func splitString() {
	res := kyu_6.Solution2("abc")
	fmt.Println(res)
}

func isValidCoordinates() {
	invalidCoordinates := []string{
		"23.234, - 23.4234",
		"2342.43536, 34.324236",
		"N23.43345, E32.6457",
		"99.234, 12.324",
		"6.325624, 43.34345.345",
		"0, 1,2",
		"0.342q0832, 1.2324",
		"23.245, 1e1"}

	for _, v := range invalidCoordinates {
		res := kyu_6.IsValidCoordinates(v)
		fmt.Println(res)
	}
}

func race() {
	res := kyu_6.Race(720, 850, 70)
	fmt.Println(res)
}

func test_6_30() {
	res := practice.Test_6_30("aaabbbbcca")
	fmt.Println(res)
}

func thirt() {
	res := kyu_6.Thirt(85299258)
	fmt.Println(res)
}

func operArray() {
	var data = []int{18, 69, -90, -78, 65, 40}
	res := kyu_6.OperArray(kyu_6.Lcmu, data, data[0])
	fmt.Println(res)
}

func main() {
	// operArray()
	// thirt()
	// test_6_30()
	// race()
	// isValidCoordinates()
	// splitString()
	// bandNameGenerator()
	// hasUniqueChar()
	// decode()
	// toNato()
	// longestConsec()
	// reverses()
	// inAscOrder()
	// duplicateCount()
	// accum()
	// isTriangle()
	// century()
	// repeatStr()
	// removeChar()
	// positiveSum()
	// evenOrOdd()
	// multiple3And5()
	// twoOldestAges()
	// toCamelCase()
	// findUniq()
	// sumEvenFibonacci()
	// validParentheses()
}
