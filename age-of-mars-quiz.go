package main

import "fmt"

const aYearOnEarth int = 365
const aYearOnMars int = 687

func main() {
    var age int
    fmt.Scanln(&age)

    mars := mars_age(age)
    fmt.Println(mars)
}

func mars_age(age int) int {
	yearOnDayEarth := age * aYearOnEarth
	return yearOnDayEarth / aYearOnMars
}