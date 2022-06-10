package main

import (
	"fmt"
	"sort"
)

func DragonOfLoowater(dragonHead, knightHeight []int) {
	solve := func(dragons, knights []int) int {
		sort.Ints(knights)
		sort.Ints(dragons)
		var total int
		for _, dragon := range dragons {
			l := len(knights)
			if pos := sort.Search(l, func(i int) bool { return knights[i] >= dragon }); pos < l {
				total += knights[pos]
				knights = knights[pos+1:]
			} else {
				return -1
			}
		}
		return total
	}

	if solve(dragonHead, knightHeight) == -1 {
		fmt.Println("knight fall")
	} else {
		fmt.Println(solve(dragonHead, knightHeight))
	}
}

func main() {
	DragonOfLoowater([]int{5, 4}, []int{7, 8, 4})		// 11
	DragonOfLoowater([]int{5, 10}, []int{5})			// knight fall
	DragonOfLoowater([]int{7, 2}, []int{4, 3, 1, 2})	// knight fall
	DragonOfLoowater([]int{7, 2}, []int{2, 1, 8, 5})	// 10
}
