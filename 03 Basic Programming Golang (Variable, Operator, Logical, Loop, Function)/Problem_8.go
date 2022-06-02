package main

func cetakTabelPerkalian(number int) {
	for i:=1; i<=number; i++ {
		for j:=1; j<=number; j++ {
			res   := i * j

			if (res < 10) {
				print("   ", res)
			} else if (res >= 10) && (res <= 99) {
				print("  ", res)
			} else if (res >= 100) {
				print(" ", res)
			}
		}
		println()
	}
}

func main() {
	cetakTabelPerkalian(9)
}