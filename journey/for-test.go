package main

import "fmt"

func main() {
	res := 0

	for i := 0; i < 10; i += 3 {
		res += i
	}

	fmt.Println(res)
}