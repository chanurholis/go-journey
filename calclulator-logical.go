package main

import "fmt"

func main() {
	var bilanganPertama, bilanganKedua float32 
	var operator string

	fmt.Scanln(&bilanganPertama)
	fmt.Scanln(&operator)
	fmt.Scanln(&bilanganKedua)

	switch {
		case operator == "+":
			tambah(bilanganPertama, bilanganKedua)
		case operator == "-":
			kurang(bilanganPertama, bilanganKedua)
		case operator == "*":
			kali(bilanganPertama, bilanganKedua)
		case operator == "/":
			bagi(bilanganPertama, bilanganKedua)
		// case operator == "%":
		// 	modulus(bilanganPertama, bilanganKedua)
		// case operator == "**":
		// 	modulus(bilanganPertama, bilanganKedua)
		default:
			fmt.Println("Opeator harus diisi!")
	}
}

func kali(bilanganPertama, bilanganKedua float32) {
	fmt.Println(bilanganPertama * bilanganKedua)
}

func bagi(bilanganPertama, bilanganKedua float32) {
	fmt.Println(bilanganPertama / bilanganKedua)
}

func kurang(bilanganPertama, bilanganKedua float32) {
	fmt.Println(bilanganPertama - bilanganKedua)
}

func tambah(bilanganPertama, bilanganKedua float32) {
	fmt.Println(bilanganPertama + bilanganKedua)
}

// func modulus(bilanganPertama, bilanganKedua float32) {
// 	fmt.Println(bilanganPertama % bilanganKedua)
// }

// func kuadrat(bilanganPertama, bilanganKedua float32) {
// 	fmt.Println(bilanganPertama ** bilanganKedua)
// }