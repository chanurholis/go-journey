package main

import "fmt"

func main() {
	x := 2

	switch {
		case x == 1:
			fmt.Println("One")
		case x == 2:
			fmt.Println("Two")
		default:
			fmt.Println("Other")
	}
}