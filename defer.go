package main

import "fmt"

func main() {
	defer exampleUseDefer()
	fmt.Println("Hey you")
}

func exampleUseDefer() {
	for i := 1; i <= 1000; i++ {
		fmt.Println(i)
	}
}