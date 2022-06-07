package main
import (
	"fmt"
	"unicode"
)

func caesar(offset int, input string) string {
	var hasil string

	for i := 0; i<len(input); i++ {
		if isUpper(input) {
			hasil += string(int(int(input[i])+offset-65)%26 + 65)
		} else {
			hasil += string(int(int(input[i])+offset-97)%26 + 97)
		}
	}
	return hasil
}

func isUpper(s string) bool {
    for _, r := range s {
        if !unicode.IsUpper(r) && unicode.IsLetter(r) {
            return false
        }
    }
    return true
}

func main() {
	fmt.Println(caesar(3, "abc"))
	fmt.Println(caesar(2, "alta"))
	fmt.Println(caesar(10, "alterracademy"))
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz"))	
}