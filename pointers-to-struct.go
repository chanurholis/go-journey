package main

import "fmt"

func main() {
	type Contact struct {
		name string
		age  int
	}

	x := Contact{"Chacha Nurholis", 20}

	y := &x

	fmt.Println("My", y.name, "age", y.age)
}
