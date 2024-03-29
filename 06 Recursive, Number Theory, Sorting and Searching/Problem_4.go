package main
import "fmt"

func MaxSequence(arr []int) int {
    var kiri  = arr[0]
    var kanan = arr[0]
    var n int = len(arr)

    for i:=1; i<n; i++ {
        kanan = maxInt(arr[i], kanan + arr[i])
        kiri = maxInt(kiri, kanan)
    }
    
    return int(kiri)
}

func maxInt(a, b int) int {
    if a < b {
        return b
    }
    return a
}

func main() {
	fmt.Println(MaxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))		// 6
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6}))		// 7
	fmt.Println(MaxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3}))		// 7
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6}))		// 8
	fmt.Println(MaxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6}))			// 12
}