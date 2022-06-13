func rotate(nums []int, k int) {
    n := len(nums)
	k %= n

	reversed(nums, 0, n-1)
	reversed(nums, 0, k-1)
	reversed(nums, k, n-1)
}

func reversed(nums []int, start int, end int) {
	for i := 0; start < end; i++ {
		nums[start], nums[end] = nums[end], nums[start]
		start++ 
        end--
	}
}