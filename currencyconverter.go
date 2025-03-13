package main

import "fmt"

func main() {
	//initialize 3 variables fixed values of conversion currencies
	const currencyconverter = 126.12
	const currencyconverter1 = 1.85
	const currencyconverter2 = 120.12

	var usd float64

	// prompt user to enter usd

	fmt.Print("Enter amount is USD: ")
	fmt.Scanf("%f", &usd)

	// initialize a converter
	kes := usd * currencyconverter
	euro := usd * currencyconverter1
	jyen := usd * currencyconverter2

	// Display the converted currency
	fmt.Printf("%.2f USD is: \n KES %.2f\n Euro %.2f\n Japanes Yen %.2f\n", usd, kes, euro, jyen)
}
