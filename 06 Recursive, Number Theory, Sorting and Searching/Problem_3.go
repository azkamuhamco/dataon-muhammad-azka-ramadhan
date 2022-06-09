package main

import (
	"fmt"
	"math"
	"strconv"
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

func primaSegiEmpat(high, wide, start int) {
	nArray := high * wide
	var arr []int

	for {
		if isPrime(start) {
			arr = append(arr, start)
		}
		start++
		if len(arr) >= nArray {
			break
		}
	}

	sum := 0
	for i:=0; i<len(arr); i++ {
	   sum += arr[i]
	}

	for j:=0; j<wide; j++ {
		for i:=0; i<high; i++ {
			a := (high * j) + i
			fmt.Printf(strconv.Itoa(arr[a]) + " ")
		}
		fmt.Println()
	}

	fmt.Println(sum)
}

func main() {
	primaSegiEmpat(2, 3, 13)
	primaSegiEmpat(5, 2, 1)
}