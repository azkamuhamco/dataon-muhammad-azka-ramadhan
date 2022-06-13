func search(nums []int, target int) int {
    l := 0
    r := len(nums) - 1
    
    for l <= r {
        tengah := l + (r - l) / 2
        if nums[tengah] == target {
            return tengah
        } else if nums[tengah] < target {
            l = tengah + 1
        } else {
            r = tengah - 1
        }
    }
    return -1
}