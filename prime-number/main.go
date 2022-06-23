package main

import(
	"fmt"
)

func main() {

	var choise int

// we choose whether we want to check one number or several at once within the set limits

	fmt.Println("enter number 1 or 2 for make your choise")
	fmt.Scanln(& choise)

	var numToCheck int
	var arr[]int

// initialize the number or numbers according to our choice

	switch {
	case choise == 1:
		fmt.Println("enter the number to check from the keyboard")
		fmt.Scanln(& numToCheck)
		arr = append(arr, numToCheck)
	case choise == 2:
		var from, to int
		fmt.Println("set boundaries\nfrom")
		fmt.Scanln(& from)
		fmt.Println("to")
		fmt.Scanln(& to)
		numToCheck = from
		for i:=from; i<=to; i++ {
			arr = append(arr, numToCheck)
			numToCheck += 1
		}
	}

// check if the numbers are prime

	for j := 0; j < len(arr); j++ {

		fmt.Printf("number %v is prime? - ", arr[j])

		if arr[j] == 1 {
			fmt.Println("NO")
			continue
		}

		if arr[j] <= 3 {
			fmt.Println("YES")
			continue
		}

		if arr[j] % 2 == 0 {
			fmt.Println("NO")
			continue
		}

		check := 3

		for step := 0; step < arr[j]; step++ {

			if check >= arr[j] {
				fmt.Println("YES")
				break
			}

			if arr[j] % check == 0 {
				fmt.Println("NO")
				break
			}

			check += 2

		}
	}
}