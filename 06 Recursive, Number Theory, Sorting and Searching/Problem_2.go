package main

import "fmt"

func fibonacci(number int) int {
	var m = []int{0, 1, 1}
	if len(m) > number || number == 0 {
		return m[number]
	}
	res := fibonacci(number-1) + fibonacci(number-2)
	m = append(m, res)
	return res
}

func main() {
	fmt.Println(fibonacci(0))		// 0
	fmt.Println(fibonacci(2))		// 1
	fmt.Println(fibonacci(9))		// 34
	fmt.Println(fibonacci(10))		// 55
	fmt.Println(fibonacci(12))		// 144
}