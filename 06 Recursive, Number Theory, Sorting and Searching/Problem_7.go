package main
import "fmt"

func PlayingDomino(cards [][]int, deck []int) interface{} {
	return Deck(deck)
}

func Deck(deck []int) []int{
	return deck
}

func main() {
	fmt.Println([][]int{[]int{6, 5}, []int{3, 4}, []int{2, 1}, []int{3, 3}}, []int{4, 3})
	fmt.Println([][]int{[]int{6, 5}, []int{3, 3}, []int{3, 4}, []int{2, 1}}, []int{3, 6})
	fmt.Println([][]int{[]int{6, 6}, []int{2, 4}, []int{3, 6}}, []int{5, 1})
}