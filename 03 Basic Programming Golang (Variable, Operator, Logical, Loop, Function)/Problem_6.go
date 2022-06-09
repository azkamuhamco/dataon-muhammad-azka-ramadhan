package main

import "fmt"

func pangkat(x, y int) int {
    if y == 0 {
		return 1
	} else if x == 0 {
		return 0
	}
    return x * pangkat(x, y - 1)
}

func main() {
	fmt.Println(pangkat(2, 3))
	fmt.Println(pangkat(5, 3))
	fmt.Println(pangkat(10, 2))
	fmt.Println(pangkat(2, 5))
	fmt.Println(pangkat(7, 3))
}