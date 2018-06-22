package main

import (
	"fmt"
	"gopractice/practice"
)

func validParentheses() {
	// var testString string
	// fmt.Scanln(&testString)
	// fmt.Println("InputSring: ", testString)

	res := practice.StackValidParentheses("{}[]<>(")
	fmt.Println(res)
}

func sumEvenFibonacci() {
	res := practice.SumEvenFibonacci(1)
	fmt.Println(res)
}

func findUniq() {
	s := []float32{1, 2, 1}
	res := practice.FindUniq(s)
	fmt.Println(res)
}

func toCamelCase() {
	s := "Th-s"
	res := practice.ToCamelCase(s)
	fmt.Println(res)
}

func twoOldestAges() {
	s := []int{6, 5, 83, 5, 3, 18}
	res := practice.TwoOldestAges(s)
	fmt.Println(res)
}

func multiple3And5() {
	res := practice.Multiple3And5(10)
	fmt.Println(res)
}

func evenOrOdd() {
	res := practice.EvenOrOdd(2)
	fmt.Println(res)
}

func positiveSum() {
	res := practice.PositiveSum([]int{1, -2, 3, 4, 5})
	fmt.Println(res)
}

func removeChar() {
	res := practice.RemoveChar("person")
	fmt.Println(res)
}

func repeatStr() {
	res := practice.RepeatStr(4, "hello ")
	fmt.Println(res)
}

func combat() {
	res := practice.Combat(100.0, 50.0)
	fmt.Println(res)
}

func century() {
	res := practice.Century(89)
	fmt.Println(res)
}

func isTriangle() {
	res := practice.IsTriangle(2, 5, 1)
	fmt.Println(res)
}

func accum() {
	res := practice.Accum("RqaEzty")
	fmt.Println(res)
}
func duplicateCount() {
	res := practice.DuplicateCount("inddivisibility")
	fmt.Println(res)
}

func inAscOrder() {
	res := practice.InAscOrder([]int{1, 2, 4, 7, 19})
	fmt.Println(res)
}

func reverses() {
	res := practice.Reverses("hello")
	fmt.Println(res)
}

func main() {
	// reverses()
	// inAscOrder()
	// duplicateCount()
	// accum()
	// isTriangle()
	// century()
	// combat()
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
