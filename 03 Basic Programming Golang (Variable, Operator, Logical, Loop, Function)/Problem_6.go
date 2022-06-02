package main

import "fmt"

func pangkat(base, pangka int) int {
	if pangka > 1 {
		return base * pangkat(base, pangka - 1)
	}
	return base
}

func main() {
	fmt.Println(pangkat(2, 3))
	fmt.Println(pangkat(5, 3))
	fmt.Println(pangkat(10, 2))
	fmt.Println(pangkat(2, 5))
	fmt.Println(pangkat(7, 3))
}