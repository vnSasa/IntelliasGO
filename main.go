package main

import "fmt"

const (
	PRICE_APPLE = 5.99
	PRICE_PEAR float32 = 7
	MY_MONEY float32 = 23
)

func main()  {
	
// TASK 1
	
	amountApples := 9
	amountPears := 8
	
	var buyApples = float32(amountApples)*PRICE_APPLE
	var buyPears = float32(amountPears)*PRICE_PEAR

	var howMuchMoney = buyApples+buyPears
	fmt.Printf("\nHow much money do you have to spend to buy %x apples and %x pears?", amountApples, amountPears)
	fmt.Println(" ---", howMuchMoney, "UAH")
	
// TASK 2

	var howMuchPears = MY_MONEY/PRICE_PEAR
	fmt.Println("How many pears can we buy? ---", int(howMuchPears), "pears")

// TASK 3

	var howMuchApples = MY_MONEY/PRICE_APPLE
	fmt.Println("How many apples can we buy? ---", int(howMuchApples), "apples")

// TASK 4

	amountApples = 2
	amountPears = 2
	var ourCapabilities bool = MY_MONEY >= float32(amountApples)*PRICE_APPLE + float32(amountPears)*PRICE_PEAR
	fmt.Printf("Can we buy %x pears and %x apples?", amountApples, amountPears)
	fmt.Println(" ---", ourCapabilities, "\n")

}