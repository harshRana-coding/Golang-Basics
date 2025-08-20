package main

import (
	"fmt"
	"example.com/investment-calculator/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"


func bank() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile,1000)

	if err != nil {
		fmt.Println("ERROR!!")
		fmt.Println(err)
		fmt.Println("-------------")
		// panic("Can't Continue, Sorry.")
	}

	fmt.Println("Welcome to the bank !")
	fmt.Println(randomdata.PhoneNumber())

	for {
		presentOptions();

		var choice int
		fmt.Print("Enter your choice :")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is", accountBalance)
		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount must be greater than zero.")
				continue
			}

			accountBalance += depositAmount
			fileops.WriteFloatToFile(accountBalanceFile, accountBalance)
			fmt.Println("Balance Updated ! New Amount: ", accountBalance)
		case 3:

			fmt.Print("Your withdrawn: ")
			var withdrawnAmount float64
			fmt.Scan(&withdrawnAmount)
			if withdrawnAmount <= 0 {
				fmt.Println("Invalid amount must be greater than zero.")
				continue
			}
			if withdrawnAmount > accountBalance {
				fmt.Println("Insufficient balance !")
				continue
			}
			accountBalance -= withdrawnAmount
			fileops.WriteFloatToFile(accountBalanceFile, accountBalance)
			fmt.Println("Balance Updated ! New Amount: ", accountBalance)
		default:
			fmt.Println("GoodBye !")
			fmt.Println("Thanks for choosing our bank")
			return
		}
	}

}
