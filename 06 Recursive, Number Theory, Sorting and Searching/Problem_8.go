// BELUM TUNTAS
package main

import (
	"fmt"
	"sort"
)

type pair struct {
	name string
	count int
}

func MostAppearItem(items []string) []pair {
	// Langkah 1: Buat mapping nama dan jumlah frekuensi
    basket := make(map[string]int)
    for _, freq :=  range items {
        basket[freq]++
    }

	// Langkah 2: Pisahkan menjadi array masing2
	keys   := make([]string, 0, len(basket))
	values := make([]int, 0, len(basket))
	for key, value := range basket {
		keys = append(keys, key)
		values = append(values, value)
	}

	// Langkah 2: Urutkan dari frekuensi dari terkecil ke terbesar
	sort.SliceStable(keys, func(i, j int) bool {
		return basket[keys[i]] < basket[keys[j]]
	})
	sort.Ints(values[:])

	// Pengembalian nilai untuk ditampilkan
	var tagsList []pair
	for i:=0; i<len(keys); i++ {
		tagsList = append(tagsList, pair { keys[i], values[i] }, )
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
