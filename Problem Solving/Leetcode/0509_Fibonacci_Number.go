var m = []int{0, 1, 1}
func fib(n int) int {
	if len(m) > n || n == 0 {
		return m[n]
	}
	res := fib(n-1) + fib(n-2)
	m = append(m, res)
	return res
}