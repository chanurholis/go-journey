package main

import "fmt"

func main() {
	fmt.Println("Start\n")

	for i := 1; i <= 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("End\n")
}