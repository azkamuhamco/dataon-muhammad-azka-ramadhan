// BELUM TUNTAS
package main

import "fmt"

type pair struct {
	name string
	count int
}

func MostAppearItem(items []string) []pair {
	var tagsList []pair
	
	for i:=0; i<len(items); i++ {
		tagsList = append(tagsList, pair { items[i] +"->", i}, )
	}

	return tagsList
}

func main() {
	fmt.Println(MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}))
	// golang->1 ruby->2 js-4
	fmt.Println(MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}))	
	// C->1 D->2 B->3 A->4
	fmt.Println(MostAppearItem([]string{"football", "basketball", "tenis"}))	
	// football->1 basketball->1 tenis->1
}
