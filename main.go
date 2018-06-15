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

func main() {
	sumEvenFibonacci()
	// validParentheses()
}
