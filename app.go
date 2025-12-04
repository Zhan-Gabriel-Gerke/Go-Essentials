package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var accountBalance = getBalanceFromFile()

	//for i := 0; i < 5; i++ {
	//	fmt.Println("Test")
	//}
	//
	//for {
	//	fmt.Println("Test")
	//}
	fmt.Println("Welcome to Go bank")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check the balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3 . Withdraw money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	//wantsCheckBalance := choice == 1

	//switch choice {
	//case 1:
	//	fmt.Println("Your balance is ", accountBalance)
	//case 2:
	//	fmt.Print("Your deposit: ")
	//default:
	//	fmt.Println("Default")
	//
	//}

	if choice == 1 {
		fmt.Println("Your balance is ", accountBalance)
	} else if choice == 2 {
		fmt.Print("Your deposit: ")
		var depositAmount float64
		fmt.Scan(&depositAmount)
		accountBalance += depositAmount
		writeBalanceToFile(accountBalance)
		fmt.Println("Balance updated! New amount: ", accountBalance)
	} else if choice == 3 {
		fmt.Print("Your withdraw: ")
		var withdrawAmount float64
		fmt.Scan(&withdrawAmount)
		if withdrawAmount > accountBalance {
			fmt.Println("Insufficient funds")
		} else {
			accountBalance -= withdrawAmount
			fmt.Println("Balance updated! New amount: ", accountBalance)
		}
	} else {
		fmt.Println("Exiting...")
	}
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() float64 {
	data, _ := os.ReadFile(accountBalanceFile)
	balanceText := string(data)
	balance, _ := strconv.ParseFloat(balanceText, 64)
	return balance
}
