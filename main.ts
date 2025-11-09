/*
* Author: Bianca Bertinato
* Version: 1.0.0
* Date: 2025-11-09
* This program calculates the total cost of a meal including tax and tip, and divides the total among the number of people splitting the bill.
*/

// CONSTANTS
const taxRate: number = 0.13; 
const tipRate: number = 0.15;

// VARIABLES
let mealCost: number = 315.99; 
let numPeople: number = 5;     

// PROCESS
let taxAmount: number = mealCost * taxRate;           
let tipAmount: number = mealCost * tipRate;           
let totalCost: number = mealCost + taxAmount + tipAmount; 
let costPerPerson: number = totalCost / numPeople;    

// OUTPUT
console.log("The cost of the meal was: $" + mealCost);
console.log("The HST amount is: $" + taxAmount);
console.log("The tip amount is: $" + tipAmount);
console.log("The total cost is: $" + totalCost);
console.log("The cost per person is: $" + costPerPerson);

console.log("\nDone.");