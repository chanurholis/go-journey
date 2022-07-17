package main

import "fmt"

func main() {
	for i := 1; i <= 3; i++ {
		var input int 

		fmt.Scanln(&input)
	
		if input >= 0 && input <= 10 {
			switch {
				case input == 0:
					fmt.Println("Zero")
				case input == 1:
					fmt.Println("One")
				case input == 2:
					fmt.Println("Two")
				case input == 3:
					fmt.Println("Three")
				case input == 4:
					fmt.Println("Four")
				case input == 5:
					fmt.Println("Five")
				case input == 6:
					fmt.Println("Six")
				case input == 7:
					fmt.Println("Seven")
				case input == 8:
					fmt.Println("Eight")
				case input == 9:
					fmt.Println("Nine")
				case input == 10:
					fmt.Println("Ten")
				default:
					fmt.Println("Other")
			}
		}
	}
}