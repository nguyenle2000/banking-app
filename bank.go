package main

import (
	"fmt"
	"log"

	"example.com/bankingapp/operation"
)

const balanceFiles = "account.txt"

func main() {
	var accountBalance, err = operation.ReadBalanceFromFiles(balanceFiles)
	if err != nil {
		log.Fatal()
	}
	var choice int
	for {
		operation.DisplayMessage()
		fmt.Println("Please choose option to continue")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Printf("Your balance is %v", accountBalance)
			continue
		case 2:
			var deposit float64
			fmt.Printf("Enter your deposit: ")
			fmt.Scan(&deposit)

			if deposit < 0 {
				fmt.Println("Deposit must be positive")
			} else {
				accountBalance += deposit
				fmt.Printf("You have deposited %v to your account. Account balance update: %v", deposit, accountBalance)
				operation.WriteBalanceToFiles(accountBalance, balanceFiles)
				continue
			}
		case 3:
			var withdraw float64
			fmt.Printf("Enter your withdrawal: ")
			fmt.Scan(&withdraw)

			if withdraw < 0 {
				fmt.Println("Withdraw must be positive")
			} else if withdraw > accountBalance {
				fmt.Println("Bạn không thể rút tiền vượt quá mức số dư trong tài khoản")
			} else {
				accountBalance -= withdraw
				fmt.Printf("You have withdraw %v to your account. Account balance update: %v", withdraw, accountBalance)
				operation.WriteBalanceToFiles(accountBalance, balanceFiles)
				continue
			}
		default:
			fmt.Println("Thanks for choosing Go Banking")
			return
		}
		fmt.Println("Goodbye!")
	}
}
