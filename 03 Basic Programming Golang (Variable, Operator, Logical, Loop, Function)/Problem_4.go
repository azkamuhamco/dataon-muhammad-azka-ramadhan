package main

import (
	"fmt"
)

func primeNumber(number int) bool {
	for i := 2; i < number; i++ {
		if (number % i) == 0 {
			return false
		}	
	}
	return true
}

func main() {
	fmt.Println(primeNumber(11))	// true
	fmt.Println(primeNumber(13))
	fmt.Println(primeNumber(17))
	fmt.Println(primeNumber(20))	// false
	fmt.Println(primeNumber(35))
}