package main

import (
	"fmt"
	"math"
	"strconv"
)

func SimpleEquations(a, b, c int) {
	var x int
	var y int
	var z int

	if c == a * ((a+b)/a) + ((a+b)/a) {
		y = ((a+b)/a)
		x = y - (b/a)
		z = y + (b/a)
		fmt.Println(strconv.Itoa(x), strconv.Itoa(y), strconv.Itoa(z))
	} else if float64(c/a) == math.Round(math.Pow(float64(b), 1.0/3)) {
		x = c/a
		y = x
		z = x
		fmt.Println(strconv.Itoa(x), strconv.Itoa(y), strconv.Itoa(z))
	} else {
		fmt.Println("no solution")
	}
}

func main() {
	SimpleEquations(1, 2, 3)	// no solution
	SimpleEquations(6, 6, 14)	// 1 2 3
}