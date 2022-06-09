package main
import "fmt"

func moneyCoins(money int) []int {
	var deno = [11]int{1, 10, 20, 50, 100, 200, 500, 1000, 2000, 5000, 10000}
	n := len(deno)
	var ans []int

	for i := n-1; i >= 0; i-- {
		for money >= deno[i] {
			money -= deno[i]
			ans = append(ans, deno[i])
		}
	}

	return ans
}

func main() {
	fmt.Println(moneyCoins(123))
	fmt.Println(moneyCoins(432))
	fmt.Println(moneyCoins(543))
	fmt.Println(moneyCoins(7752))
	fmt.Println(moneyCoins(15321))
}