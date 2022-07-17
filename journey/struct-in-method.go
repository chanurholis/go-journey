package main

import "fmt"

type Contact struct {
	name string
	age  int
}

func main() {
	x := Contact{"Chacha Nurholis", 20}

	name := getName(x)
	age := getAge(x)

	fmt.Println("Name", name)
	fmt.Println("Age", age)
}

func getName(x Contact) string {
	return x.name
}

func getAge(x Contact) int {
	return x.age
}
