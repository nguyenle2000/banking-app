package main

import (
	"fmt"

	"example.com/bankingapp/operation"
)

const accountBalanceFiles = "balance.txt"

func main() {
	var userAccountBalance = 0.0
	userAccountBalance = operation.ReadFiles(accountBalanceFiles)
	for {
		operation.DisplayMessage()
		fmt.Println("Please choose your option to continue")
		optioninput := operation.PromptInput("Option: ")

		switch optioninput {
		case 1:
			fmt.Printf("Your account balance: %v\n", userAccountBalance)
			continue
		case 2:
			var userDepositAmount float64
			fmt.Printf("Please enter your deposit amount\n")
			fmt.Scan(&userDepositAmount)

			if userDepositAmount <= 0 {
				fmt.Printf("Your deposit amount should be a positive number. Please try again!\n")
			} else {
				userAccountBalance += userDepositAmount
				fmt.Printf("You've successfully deposit %v to your bank account\n Account Balance is updated: %v\n", userDepositAmount, userAccountBalance)
				operation.WriteFiles(userAccountBalance, accountBalanceFiles)
				continue
			}
		case 3:
			var userWithdrawAmount float64
			fmt.Printf("Please enter your deposit amount: ")
			fmt.Scan(&userWithdrawAmount)

			if userWithdrawAmount <= 0 {
				fmt.Printf("Your withdrawal amount should be a positive number. Please try again!\n")
			} else if userWithdrawAmount > userAccountBalance {
				fmt.Println("Withdrawal Amount cannot exceeds your account balance")
			} else {
				userAccountBalance -= userWithdrawAmount
				fmt.Printf("You've successfully deposit %v to your bank account\n Account Balance is updated: %v\n", userWithdrawAmount, userAccountBalance)
				operation.WriteFiles(userAccountBalance, accountBalanceFiles)
				continue
			}
		case 4:
			fmt.Println("Calculate Earning After Tax")
			var revenue float64
			var expenses float64
			var taxRate float64
			fmt.Printf("Please input your profit here: ")
			fmt.Scan(&revenue)
			fmt.Printf("Please input your expenses: ")
			fmt.Scan(&expenses)
			fmt.Printf("Please Input your tax rate: ")
			fmt.Scan(&taxRate)

			beforeTaxEarning := operation.EarningBeforeTax(revenue, expenses)
			fmt.Printf("Your earning before tax: %v\n", beforeTaxEarning)

			earningAfterTax := operation.EarningAfterTax(beforeTaxEarning, taxRate)

			fmt.Printf("Your earning after tax: %v\n", earningAfterTax)
			continue

		default:
			fmt.Println("Thanks for choosing Go Bank")
		}

	}
}
