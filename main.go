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

func main() {
	evenOrOdd()
	// multiple3And5()
	// twoOldestAges()
	// toCamelCase()
	// findUniq()
	// sumEvenFibonacci()
	// validParentheses()
}
