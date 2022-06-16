package main

import "fmt"

const (
	PRICE_APPLE float32 = 5.99
	PRICE_PEAR float32 = 7
	MY_MONEY float32 = 23
)

func main()  {
	
	var howMuchMoney float32 = 9*PRICE_APPLE+8*PRICE_PEAR 
	fmt.Println("\nHow much money do you have to spend to buy 9 apples and 8 pears? ---", howMuchMoney, "UAH")
	
	var howMuchPears = MY_MONEY/PRICE_PEAR
	fmt.Println("How many pears can we buy? ---", int(howMuchPears), "pears")

	var howMuchApples = MY_MONEY/PRICE_APPLE
	fmt.Println("How many apples can we buy? ---", int(howMuchApples), "apples")

	var ourCapabilities bool = MY_MONEY >= 2*PRICE_APPLE+2*PRICE_PEAR
	fmt.Println("Can we buy 2 pears and 2 apples? ---", ourCapabilities, "\n")

}