package main
import "fmt"

func BinarySearch(array []int, x int) {
    kembali := func(nums []int, target int) int {
        start := 0
        end   := len(nums) - 1
    
        for start <= end {
            mid := (end + start) >> 1
            if nums[mid] == target {
                return mid
            } else if nums[mid] < target {
                start = mid + 1
            } else {
                end = mid - 1
            }
        }
        return -1        
    }
	fmt.Println(kembali(array, x))
}

func main() {
	BinarySearch([]int{1, 1, 3, 5, 5, 6, 7}, 3)						// 2
	BinarySearch([]int{1, 2, 3, 5, 6, 8, 10}, 5)					// 3
	BinarySearch([]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 53)		// 6
	BinarySearch([]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 100)	// -1
}