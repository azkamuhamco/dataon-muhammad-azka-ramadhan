func minMaxGame(nums []int) int {
    for len(nums) != 1 {
        tmp := []int{}
        for i := 0; i < len(nums) / 2; i++ {
            if i & 1 == 0 {
                tmp = append(tmp, min(nums[2 * i], nums[2 * i + 1]))
            } else {
                tmp = append(tmp, max(nums[2 * i], nums[2 * i + 1]))
            }
        }
        nums = tmp
    }
    return nums[0]
}

func min(a, b int) int {
    if a < b { return a }
    return b
}

func max(a, b int) int {
    if a > b { return a }
    return b
}