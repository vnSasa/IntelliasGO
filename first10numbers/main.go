package main

import (
	"fmt"
)

func main() {

	// you can initialize any number here
	num := 11

	// set the counter of the displayed numbers
	outputedNum := 0

	// determine the initial value of the variable to be checked
	check := num

	// implement a loop that will run until we output the first 10 numbers
	for outputedNum < 10 {
		
		if check % num == 0 {
			fmt.Println(check)
			outputedNum += 1
		}

		check += 1

	}
}