package main
import "fmt"

func playingDomino(cards [][]int, deck []int) interface{} {
	// Tahap 1: Temukan kartu yang salah satu angkanya sama dengan deck
	var arr [][]int
    for i:=0; i<len(cards); i++ {
		if deck[0] == cards[i][0] || deck[1] == cards[i][0] || deck[0] == cards[i][1] || deck[1] == cards[i][1] {
			arr = append(arr, cards[i])
		}
	}

	// Jika tidak ditemukan maka...
	if len(arr) == 0 {
		return "tutup kartu"
	}

	// Tahap 2: Cari kartu yang paling besar hasil penjumlahannya
	var eMax []int
	n := len(arr)

	if n == 1 {
		eMax = arr[0]
	}

	if arr[0][0] + arr[0][1] > arr[1][0] + arr[1][1] {
		eMax = arr[0]
	} else {
		eMax = arr[1]
	}

	for i:=2; i<n; i++ {
		if arr[0][0] + arr[0][1] > eMax[0] + eMax[1] { 
			eMax = arr[i]
		}
	}

	return eMax
}

func main() {
	fmt.Println(playingDomino([][]int{[]int{6, 5}, []int{3, 4}, []int{2, 1}, []int{3, 3}}, []int{4, 3}))
	fmt.Println(playingDomino([][]int{[]int{6, 5}, []int{3, 3}, []int{3, 4}, []int{2, 1}}, []int{3, 6}))
	fmt.Println(playingDomino([][]int{[]int{6, 6}, []int{2, 4}, []int{3, 6}}, []int{5, 1}))
}