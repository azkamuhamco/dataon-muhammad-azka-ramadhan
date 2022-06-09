package main

import "fmt"

func pangkat(x, y int) int {
	result := 1
	for y > 0 {
		if y%2 == 0 {
			x = x*x
			y = y>>1
		}
		result *= x
		y -= 1
	}
	return result
}

func main() {
	fmt.Println(pangkat(2, 3))
	fmt.Println(pangkat(5, 3))
	fmt.Println(pangkat(10, 2))
	fmt.Println(pangkat(2, 5))
	fmt.Println(pangkat(7, 3))
}