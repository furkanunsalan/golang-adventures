package main

import (
	"fmt"
	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountFile = "balance.txt"

func main() {

	var accountBalance, err = fileops.GetFloatFromFile(accountFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("----")
		return
		// panic(err)
	}

	fmt.Println("Welcome to GoBank")
	fmt.Println(randomdata.PhoneNumber())

	for {
		presentOptions()

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
			fileops.WriteFloatToFile(accountBalance, accountFile)
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
			fileops.WriteFloatToFile(accountBalance, accountFile)
			fmt.Println("Balance updated, new balance: ", accountBalance)
		} else {
			fmt.Println("Goodbye!")
			// return
			break
		}
	}

	fmt.Println("Thanks for choosing our bank")

}
