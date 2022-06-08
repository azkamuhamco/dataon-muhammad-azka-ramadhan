package main
import "fmt"

func pow(x, n int) int {
	if n == 0 {
		return 1
	} else if n > 0 {
		if n % 2 == 0 {
			return int(pow((x*x), int(n/2)))
		} else {
			return int(pow((x*x), int((n-1)/2))) * x
		}
	}
	return int(pow(1/x, -n))
}

func main() {
   fmt.Println(pow(2, 3))  // 8
   fmt.Println(pow(5, 3))  // 125
   fmt.Println(pow(10, 2)) // 100
   fmt.Println(pow(2, 5))  // 32
   fmt.Println(pow(7, 3))  // 343
}
