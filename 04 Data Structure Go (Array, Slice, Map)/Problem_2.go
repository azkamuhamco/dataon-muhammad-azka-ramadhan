package main

import (
	"fmt"
	"strconv"
)

func munculSekali(angka string) []int {
	var a string
	var c []int

	for i:=0; i<len(angka); i++ {
		flag := 0
		for j:=0; j<len(angka); j++ {
			if string(angka[i]) == string(angka[j]) && i!=j {
				flag = 1
				break
			}
		}

		if flag == 0 {
			a += string(angka[i])
		}
	}

    b := []rune(a)
	for i:=0; i<len(b); i++ {
		bi, _ := strconv.Atoi(string(b[i]))
		c = append(c, bi)
	}

	return c
}

func main() {
	fmt.Println(munculSekali("1234123"))		// [4]
	fmt.Println(munculSekali("76523752"))		// [6 3]
	fmt.Println(munculSekali("12345"))			// [1 2 3 4 5]
	fmt.Println(munculSekali("1122334455"))		// []
	fmt.Println(munculSekali("0872504"))		// [8 7 2 5 4]
}
