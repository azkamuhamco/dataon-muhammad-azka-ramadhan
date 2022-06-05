package main
import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	// Inisiasi variabel
	keys   := make(map[string]bool)
	input  := append(arrayA, arrayB...)
	output := []string{}
	
	// Logic untuk remove duplikat dan menambahkan petik dua beserta koma
    for i, element := range input {
        if _, value := keys[element]; !value {
            keys[element] = true

			if i != len(input) - 1 {
				output = append(output, "\"" + element + "\",")
			} else {
				output = append(output, "\"" + element + "\"")
			}
        }
    }
	return output
}

func main() {
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	fmt.Println(ArrayMerge([]string{}, []string{}))
}
