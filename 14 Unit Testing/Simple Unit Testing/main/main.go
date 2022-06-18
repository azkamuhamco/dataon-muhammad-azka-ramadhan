package main

import "fmt"

func Addition(a, b int) int {
	return a + b
}

func Subtraction(a, b int) int {
	return a - b
}

func Division(a, b int) int {
	return a / b
}

func Multiplication(a, b int) int {
	return a * b
}

func main() {
	fmt.Println(Addition(6, 2))       // 8
	fmt.Println(Subtraction(6, 2))    // 4
	fmt.Println(Division(6, 2))       // 3
	fmt.Println(Multiplication(6, 2)) // 12
}
