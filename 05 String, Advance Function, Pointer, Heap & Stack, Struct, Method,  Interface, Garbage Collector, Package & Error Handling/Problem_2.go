package main
import "fmt"

func caesar(offset int, input string) string {
	var hasil string

	for i := 0; i<len(input); i++ {
		hasil += string(int(int(input[i])+offset-97)%26 + 97)
	}
	return hasil
}

func main() {
	fmt.Println(caesar(3, "abc"))
	fmt.Println(caesar(2, "alta"))
	fmt.Println(caesar(10, "alterracademy"))
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz"))	
}