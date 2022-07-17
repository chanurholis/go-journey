package main

import(
	"fmt"
	"reflect"
)

// Pointers are special variables that hold the memory address of values.

var p *int

func main() {
	fmt.Println(reflect.TypeOf(p))
	fmt.Println(p)
}