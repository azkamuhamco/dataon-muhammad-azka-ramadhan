package main
import "fmt"

func getMinMax(numbers ...*int) (min int, max int) {
	var mmi int
	var mma int
	n := len(numbers)

    if n == 1 {
		mma = *numbers[0]
		mmi = *numbers[0]	
	}

	if *numbers[0] > *numbers[1] {
		mma = *numbers[0]
		mmi = *numbers[1]
	} else {
		mma = *numbers[1]
		mmi = *numbers[0]
	}

	for i:=2; i<n; i++ {
		if *numbers[i] > mma {
			mma = *numbers[i]
		} else if *numbers[i] < mmi {
			mmi = *numbers[i]
		}
	}

	return mmi, mma
}

func main() {
	var a1, a2, a3, a4, a5, a6, min, max int
	fmt.Scan(&a1)
	fmt.Scan(&a2)
	fmt.Scan(&a3)
	fmt.Scan(&a4)
	fmt.Scan(&a5)
	fmt.Scan(&a6)
	min, max = getMinMax(&a1, &a2, &a3, &a4, &a5, &a6)
	fmt.Println(max, "is the maximum number")	
	fmt.Println(min, "is the minimum number")	
}