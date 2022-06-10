// MASIH SALAH
package main
import "fmt"

func playingDomino(cards [][]int, deck []int) interface{} {
	var hasil []int

    for i:=0; i<len(cards); i++ {
		for j:=0; j<2; j++ {
			if deck[0] == cards[i][j] || deck[1] == cards[i][j]{
				hasil = append(hasil, cards[i]...)
			}
		}
	}

	if len(hasil) == 0 {
		return "tutup kartu"
	}
	return hasil
}

func main() {
	fmt.Println(playingDomino([][]int{[]int{6, 5}, []int{3, 4}, []int{2, 1}, []int{3, 3}}, []int{4, 3}))
	fmt.Println(playingDomino([][]int{[]int{6, 5}, []int{3, 3}, []int{3, 4}, []int{2, 1}}, []int{3, 6}))
	fmt.Println(playingDomino([][]int{[]int{6, 6}, []int{2, 4}, []int{3, 6}}, []int{5, 1}))
}