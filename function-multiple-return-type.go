package main

import "fmt"

func main() {
	a, b := multipleReturnType(1, 2)

	fmt.Println(a, b)
}

// Multiple return type int int
func multipleReturnType(bilanganSatu, bilanganDua int) (int, int) {
	return bilanganSatu, bilanganDua
}