package main

import (
	"fmt"
	"strconv"
)

func FindMinAndMax(arr []int) string {
	var nMax int
	var iMax int
	var nMin int
	var iMin int
	n := len(arr)

	if n == 1 {
		nMax = arr[0]
		nMin = arr[0]
		iMax = 0
		iMin = 0
	}

	if arr[0] > arr[1] {
		nMax = arr[0]
		iMax = 0
		nMin = arr[1]
		iMin = 1
	} else {
		nMax = arr[1]
		iMax = 1
		nMin = arr[0]
		iMin = 0
	}

	for i:=2; i<n; i++ {
		if arr[i] > nMax {
			nMax = arr[i]
			iMax = i
		} else if arr[i] < nMin {
			nMin = arr[i]
			iMin = i
		}
	}

	return "min: " + strconv.Itoa(nMin) + " index: " + strconv.Itoa(iMin) + " max: " + 
			strconv.Itoa(nMax) + " index: " + strconv.Itoa(iMax)
}

func main() {
	fmt.Println(FindMinAndMax([]int{5, 7, 4, -2, -1, 8}))
	fmt.Println(FindMinAndMax([]int{2, -5, -4, 22, 7, 7}))
	fmt.Println(FindMinAndMax([]int{4, 3, 9, 4, -21, 7}))
	fmt.Println(FindMinAndMax([]int{-1, 5, 6, 4, 2, 18}))
	fmt.Println(FindMinAndMax([]int{-2, 5, -7, 4, 7, -20}))	
}