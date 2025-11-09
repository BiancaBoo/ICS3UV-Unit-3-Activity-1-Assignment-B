/*
 * @author Bianca Bertinato
 * @version 1.0.0
 * @date 2025-10-08
 * @fileoverview This program calculates the total cost of a meal including tax and tip, and divides the total among the number of people splitting the bill.
 */
package main

import "fmt"

func main() {
	  // CONSTANTS
		const taxRate float64 = 0.13
		const tipRate float64 = 0.15

		// VARIABLES
		var mealCost float64 = 315.99
		var numPeople float64 = 5

		// PROCESS
		// Calculate tax, tip, total and per person amount
		taxAmount := mealCost * taxRate
		tipAmount := mealCost * tipRate
		totalCost := mealCost + taxAmount + tipAmount
		costPerPerson := totalCost / numPeople

		// OUTPUT
		fmt.Println("The cost of the meal was: $",mealCost)
		fmt.Println("The HST amount is: $",taxAmount)
		fmt.Println("The tip amount is: $",tipAmount) 
		fmt.Println("The total cost is: $",totalCost)
		fmt.Println("The cost per person is: $",costPerPerson)

		fmt.Println("\nDone.")
}