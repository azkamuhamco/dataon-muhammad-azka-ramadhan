package main

import (
	"fmt"
	"math"
)

func MaxSequence(arr []int) int {
    var kiri float64  = float64(arr[0])
    var kanan float64 = float64(arr[0])
    var n int         = len(arr)

    for i:=1; i<n; i++ {
        kanan = math.Max(float64(arr[i]), float64(kanan) + float64(arr[i]))
        kiri = math.Max(float64(kiri), float64(kanan))
    }
    
    return int(kiri)
}

func main() {
	fmt.Println(MaxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))		// 6
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6}))		// 7
	fmt.Println(MaxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3}))		// 7
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6}))		// 8
	fmt.Println(MaxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6}))			// 12
}