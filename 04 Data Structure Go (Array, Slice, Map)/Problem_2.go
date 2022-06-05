package main

import (
	"fmt"
	"strconv"
	"strings"
)

func munculSekali(angka string) []int {
	// inisiasi variabel
	arrae := strings.SplitAfter(angka, "")
	peta  := make(map[string]int)
	hasil := []int{}

	// jumlah frekuensi tiap angka
    for _, element := range arrae {
		peta[element]++
    }

	// Filter rune hanya freq = 1; string >>> int
	for kunci, frek := range peta {
		if frek == 1 {
			intStr, err := strconv.Atoi(kunci)
			if err != nil {
				return nil
			}
			hasil = append(hasil, intStr)
		}
	}

	return hasil
}

func main() {
	fmt.Println(munculSekali("1234123"))		// [4]
	fmt.Println(munculSekali("76523752"))		// [6 3]
	fmt.Println(munculSekali("12345"))			// [1 2 3 4 5]
	fmt.Println(munculSekali("1122334455"))		// []
	fmt.Println(munculSekali("0872504"))		// [8 7 2 5 4]
}
