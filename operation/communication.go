package operation

import "fmt"

func DisplayMessage() {
	fmt.Println("Welcome to Go Banking App")
	fmt.Println("1. Check account")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("------------------------")
	fmt.Println("4. Calculate revenue before tax")
	fmt.Println("5. Exit!")
}

func PromptInput(promptInput string) int {
	var option int
	fmt.Print(promptInput)
	fmt.Scan(&option)
	return option
}
