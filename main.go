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
	
	buyApples := float32(amountApples)*PRICE_APPLE
	buyPears := float32(amountPears)*PRICE_PEAR

	howMuchMoney := buyApples+buyPears
	fmt.Printf("\nHow much money do you have to spend to buy %x apples and %x pears?", amountApples, amountPears)
	fmt.Println(" ---", howMuchMoney, "UAH")
	
// TASK 2

	howMuchPears := MY_MONEY/PRICE_PEAR
	fmt.Println("How many pears can we buy? ---", int(howMuchPears), "pears")

// TASK 3

	howMuchApples := MY_MONEY/PRICE_APPLE
	fmt.Println("How many apples can we buy? ---", int(howMuchApples), "apples")

// TASK 4

	amountApples = 2
	amountPears = 2
	isOurCapabilities := MY_MONEY >= float32(amountApples)*PRICE_APPLE + float32(amountPears)*PRICE_PEAR
	fmt.Printf("Can we buy %x pears and %x apples?", amountApples, amountPears)
	fmt.Println(" ---", isOurCapabilities, "\n")

}