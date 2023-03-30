package main

import "fmt"

func main() {
	x := 8

	switch y := x % 2; y {
		case 0:
			x -= 1
		case 1:
			x += 1
	}

	fmt.Println(x)
}