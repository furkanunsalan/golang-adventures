package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountFile = "balance.txt"

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountFile)

	if err != nil {
		return 1000, errors.New("failed to find balance file")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("failed to parse store")
	}

	return balance, nil
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountFile, []byte(balanceText), 0644)
}

func main() {

	var accountBalance, err = getBalanceFromFile()

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("----")
		return
		// panic(err)
	}

	fmt.Println("Welcome to GoBank")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		var choice int

		fmt.Print("Your Choice: ")
		fmt.Scan(&choice)

		// switch choice {
		// case 1:
		// 	fmt.Println("Your balance: ", accountBalance)

		// case 2:
		// 	fmt.Print("Your deposit amount: ")
		// 	var depositAmount float64
		// 	fmt.Scan(&depositAmount)

		// 	if depositAmount <= 0 {
		// 		fmt.Println("İnvalid amount! Must be greater than 0.")
		// 		// return
		// 		continue
		// 	}

		// 	accountBalance += depositAmount
		// 	fmt.Println("Balance updated, new balance: ", accountBalance)

		// case 3:
		// 	fmt.Print("Your withdrawal amount: ")
		// 	var withdrawalAmount float64
		// 	fmt.Scan(&withdrawalAmount)

		// 	if withdrawalAmount <= 0 {
		// 		fmt.Println("Invalid amount! Must be greater than 0.")
		// 		continue
		// 	}

		// 	if withdrawalAmount > accountBalance {
		// 		fmt.Println("Invalid amount! You can withdraw more than you have.")
		// 		continue
		// 	}

		// 	accountBalance -= withdrawalAmount
		// 	fmt.Println("Balance updated, new balance: ", accountBalance)

		// default: // this section differs from the else case in if conditions
		// 	fmt.Println("Goodbye!")
		// 	fmt.Println("Thanks for choosing our bank") // must be moved to default rather than the outside
		// 	return
		// 	// break
		// }

		if choice == 1 {
			fmt.Println("Your balance: ", accountBalance)
		} else if choice == 2 {
			fmt.Print("Your deposit amount: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("İnvalid amount! Must be greater than 0.")
				// return
				continue
			}

			accountBalance += depositAmount
			writeBalanceToFile(accountBalance)
			fmt.Println("Balance updated, new balance: ", accountBalance)
		} else if choice == 3 {
			fmt.Print("Your withdrawal amount: ")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Println("Invalid amount! Must be greater than 0.")
				continue
			}

			if withdrawalAmount > accountBalance {
				fmt.Println("Invalid amount! You can withdraw more than you have.")
				continue
			}

			accountBalance -= withdrawalAmount
			writeBalanceToFile(accountBalance)
			fmt.Println("Balance updated, new balance: ", accountBalance)
		} else {
			fmt.Println("Goodbye!")
			// return
			break
		}
	}

	fmt.Println("Thanks for choosing our bank")

}
