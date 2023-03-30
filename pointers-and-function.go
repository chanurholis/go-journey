package main

import "fmt"

func change(val int) {
	val = 8
}

func change_ptr(ptr *int) {
	*ptr = 8
}

// When you run the code, you will see that the change() function did not change the value of our x variable, because the argument is just a copy of its value, while the change_ptr() did change the actual value of x, because it used its memory address as the argument.

func main() {
	x := 42

	change(x)
	fmt.Println(x)

	change_ptr(&x)
	fmt.Println(x)
}
