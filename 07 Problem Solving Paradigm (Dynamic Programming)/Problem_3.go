package main
import (
	"fmt"
	"unsafe"
	"math"
)

func Frog(jumps []int) int {
	n := int(unsafe.Sizeof(jumps) / unsafe.Sizeof(jumps[0]))

    // Base Case: When N < 3
    if n < 3 {
        return jumps[0];
    }
 
    // Store the results in table
    dp := make([]int, n)
 
    // Initialize base cases
    dp[0] = jumps[0];
    dp[1] = dp[0] + jumps[1] + jumps[2];

    // Iterate over the range[2, N - 2] to construct the dp array
    for i := 2; i < n - 1; i++ {
        dp[i] = int(math.Min(float64(dp[i-2] + jumps[i]), float64(dp[i-1] + jumps[i] + jumps[i+1])));
	}

	// Handle case for the last index, i.e. N - 1
    dp[n-1] = int(math.Min(float64(dp[n- 2]), float64(dp[n-3] + jumps[n-1]))) - jumps[len(jumps)-1];
 
    // Print the answer
    return dp[n-1];
}

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20}))			// 30
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50}))	// 40
}