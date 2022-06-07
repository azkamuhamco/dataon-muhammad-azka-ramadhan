package main
import "fmt"

type Student struct {
	name  []string	
	score []int
}

func (s Student) Average() float64 {
	var total float64
    for _, number := range s.score {
        total = total + float64(number)
    }
	var average float64 = total / float64(len(s.score))
	return average
}

func (s Student) Min() (min int, name string) {
	var mmi int
	var nam string
	n := len(s.score)

	if n == 1 {
		mmi = s.score[0]
		nam = s.name[0]
	}

	if s.score[0] > s.score[1] {
		mmi = s.score[1]
		nam = s.name[1]
	} else {
		mmi = s.score[0]
		nam = s.name[0]
	}

	for i:=2; i<n; i++ {
		if s.score[i] < mmi {
			mmi = s.score[i]
			nam = s.name[i]
		}
	}

	return mmi, nam
}

func (s Student) Max() (max int, name string) {
	var mma int
	var nam string
	n := len(s.score)

	if n == 1 {
		mma = s.score[0]
		nam = s.name[0]
	}

	if s.score[0] > s.score[1] {
		mma = s.score[0]
		nam = s.name[0]
	} else {
		mma = s.score[1]
		nam = s.name[1]
	}

	for i:=2; i<n; i++ {
		if s.score[i] > mma {
			mma = s.score[i]
			nam = s.name[i]
		}
	}

	return mma, nam
}

func main() {
	var a = Student{}

	for i:=0; i<6; i++ {
		var name string
		fmt.Print("Input " + string(i) + " Student's Name : ")
		fmt.Scan(&name)
		a.name = append(a.name, name)
		
		var score int
		fmt.Print("Input " + name + " Score : ")
		fmt.Scan(&score)
		a.score = append(a.score, score)
	}

	fmt.Println("\n\nAverage Score Students is", a.Average())
	
	scoreMax, nameMax := a.Max()
	fmt.Println("Max Score Students is : " + nameMax + " (", scoreMax, ")")

	scoreMin, nameMin := a.Min()
	fmt.Println("Min Score Students is : " + nameMin + " (", scoreMin, ")")
}