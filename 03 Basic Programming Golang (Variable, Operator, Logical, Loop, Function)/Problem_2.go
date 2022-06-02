package main

func main() {
	namaSiswa := "Azka"
	var studentScore int = 70
	var ket string

	if studentScore >= 80 && studentScore <= 100 {
		ket = "Nilai A"
	} else if studentScore >= 65 && studentScore <= 79 {
		ket = "Nilai B"
	} else if studentScore >= 50 && studentScore <= 64 {
		ket = "Nilai C"
	} else if studentScore >= 35 && studentScore <= 49 {
		ket = "Nilai D"
	} else if studentScore >= 0 && studentScore <= 34 {
		ket = "Nilai E"
	} else {
		ket = "Nilai invalid"
	}
	println(namaSiswa, ket)
}