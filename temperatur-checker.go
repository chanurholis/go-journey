package main

import "fmt"

func main() {
	var input float32
	const normal = 99.5

	fmt.Scanln(&input)

	if input <= normal {
		fmt.Println("Allowed")
	} else {
		fmt.Println("Fever")
	}
}