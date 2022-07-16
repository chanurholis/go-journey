package main

import "fmt"

func main() {
	result := concat("lorem", "ipsum")

	fmt.Println(result)
}

// Return type string
func concat(kataSatu, kataDua string) string {
	return kataSatu + " " + kataDua
}