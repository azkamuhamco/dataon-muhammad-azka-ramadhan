package main

import "fmt"

func main() {
	var bilangan int
	fmt.Println("Masukkan integer")
	fmt.Scanf("%d", &bilangan)

	fmt.Println("\nOutput")
	for i := 1; i <= bilangan; i++ {
		if (bilangan % i) != 0 {
			continue
		}
		fmt.Println(i)
	}
}