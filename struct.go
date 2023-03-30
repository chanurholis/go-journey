package main

import "fmt"

// Go does not support classes. Instead, it has structs.
// Structs are collections of fields that allow you to group data together.

func main() {
	type Contact struct {
		name string
		age  int
	}

	x := Contact{"Chacha Nurholis", 20}

	name := x.name
	age := x.age

	fmt.Println("Name", name)
	fmt.Println("Age", age)
}
