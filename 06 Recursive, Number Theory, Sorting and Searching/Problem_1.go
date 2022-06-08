package main

import (
	"fmt"
	"math"
)

func isPrime(number int) bool {
	numSqrt := math.Sqrt(float64(number))
	for i:=2; i<=int(numSqrt); i++ {
		if (number % i) == 0 {
			return false
		}
	}
	return true
}

func primeX(number int) int {
	var arr []int
	i:=2;
	
	for {
		if isPrime(i) {
			arr = append(arr, i)
		}
		i++
		if len(arr) >= number {
			break
		}
	}
	return arr[number-1]
}

func main() {
	fmt.Println(primeX(1))		// 2
	fmt.Println(primeX(5))		// 11
	fmt.Println(primeX(8))		// 19
	fmt.Println(primeX(9))		// 23
	fmt.Println(primeX(10))		// 29
}