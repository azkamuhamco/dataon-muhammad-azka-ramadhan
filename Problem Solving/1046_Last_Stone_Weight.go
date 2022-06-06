func lastStoneWeight(stones []int) int {
    for {
        x := 0
        y := 0
        index := 0
        
        if len(stones) == 0 {
            return 0
        }
        
        if len(stones) == 1 {
            return stones[0]
        }
        
        for i, v := range stones {
            if v > y {
                y = v
                index = i
            }
        }
        
        stones = append(stones[:index], stones[index + 1:]...)
        for i, v := range stones {
            if v > x {
                x = v
                index = i
            }
        }
        
        stones = append(stones[:index], stones[index + 1:]...)
        if y != x {
            stones = append(stones, y - x)
        }
    }
    return -1
}