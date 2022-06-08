// Fibonacci bottom-up
package main
import "fmt"

func fibo(n int) int {
    dp := []int{0, 1}
    for i := 2; i < n+1; i++ {
        dp = append(dp, (dp[i-1] + dp[i-2]))
    }
    return dp[n]
}

func main() {
	fmt.Println(fibo(0))	// 0
	fmt.Println(fibo(1))	// 1
	fmt.Println(fibo(2))	// 1
	fmt.Println(fibo(3))	// 2
	fmt.Println(fibo(5))	// 5
	fmt.Println(fibo(6))	// 8
	fmt.Println(fibo(7))	// 13
	fmt.Println(fibo(9))	// 13
	fmt.Println(fibo(10))	// 55
}