package main

import(
	"fmt"
	"net/http"
)

func main() {
	response, err := http.Get("https://chachanurholis.id")

	fmt.Println("RESPONSE", response)
	fmt.Println("\nERROR", err)
}
