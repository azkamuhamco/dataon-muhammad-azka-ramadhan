func sortedSquares(nums []int) []int {
    L   := len(nums)
    l   := 0
    r   := L-1
    ans := make([]int, L)
    
    for i:=L-1; i>=0; i-- {
        if abs(nums[l]) > abs(nums[r]) {
            ans[i] = nums[l] * nums[l]
            l++
        } else {
            ans[i] = nums[r] * nums[r]
            r--
        }
    }
    return ans
}

func abs(n int) int {
    if n > 0 { return n }
    return -n
}