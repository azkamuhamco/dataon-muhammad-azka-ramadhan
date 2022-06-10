package main

import (
	"fmt"
	"sort"
)

func DragonOfLoowater(dragonHead, knightHeight []int) {
	solve := func(dragons, knights []int) interface{} {
		sort.Ints(knights)
		sort.Ints(dragons)
		var total int
		for _, dragon := range dragons {
			l := len(knights)
			if pos := sort.Search(l, func(i int) bool { return knights[i] >= dragon }); pos < l {
				total += knights[pos]
				knights = knights[pos+1:]
			} else {
				return "knight fall"
			}
		}
		return total
	}
	fmt.Println(solve(dragonHead, knightHeight))
}

func main() {
	DragonOfLoowater([]int{5, 4}, []int{7, 8, 4})		// 11
	DragonOfLoowater([]int{5, 10}, []int{5})			// knight fall
	DragonOfLoowater([]int{7, 2}, []int{4, 3, 1, 2})	// knight fall
	DragonOfLoowater([]int{7, 2}, []int{2, 1, 8, 5})	// 10
}
