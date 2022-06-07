package main
import "fmt"

func PairSum(arr []int, target int) []int {
    var output []int
	numsMap := make(map[int]int)
	
	// convert slice to map
	for key, value := range arr {
		numsMap[value] = key
		result := target - value // find in map

		if nextKey, exist := numsMap[result]; exist && nextKey != key {
			output = append(output, nextKey, key)
			break
		}
	}

	return output
}

func main() {
	fmt.Println(PairSum([]int{1, 2, 3, 4, 6}, 6))	// [1 3]
	fmt.Println(PairSum([]int{2, 5, 9, 11}, 11))	// [0 2]
	fmt.Println(PairSum([]int{1, 3, 5, 7}, 12))		// [2 3]
	fmt.Println(PairSum([]int{1, 4, 6, 8}, 10))		// [1 2]
	fmt.Println(PairSum([]int{1, 5, 6, 7}, 6))		// [0 1]
}