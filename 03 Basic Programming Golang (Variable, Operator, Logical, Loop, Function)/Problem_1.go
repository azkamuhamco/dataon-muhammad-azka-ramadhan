package main

import (
	"fmt"
	"math"
)

func main() {
	var T float64
	var r float64
	var phi float64

	fmt.Println("Masukkan Tinggi<spasi>Luas")
	fmt.Scanf("%g %g", &T, &r)

	if math.Mod(r, 7) == 0 {
		phi = 22/7
	} else {
		phi = 3.14
	}

	var luas float64 = 2*phi*r * (r+T)
	fmt.Printf("%g", luas)
}