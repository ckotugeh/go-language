package main

import "fmt"

func main() {
	//initialize 3 variables fixed values of conversion currencies
	const currencyconverter = 129.45
	const currencyconverter1 = 0.92
	const currencyconverter2 = 147.71

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
