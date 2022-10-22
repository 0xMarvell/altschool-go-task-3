package atm

import (
	"fmt"
	"os"
	"strings"

	"github.com/0xMarvell/altschool-go-task-3/pkg/utils"
)

var acctBalance float64

// WelcomeScreen displays welcome message.
func WelcomeScreen() {
	fmt.Println("---------------ATM------------------")
	utils.NewLine(1)
	fmt.Println("\t      WELCOME!")
}

// Menu displays all available operations to user.
func Menu() {
	fmt.Println("----------SELECT OPERATION----------")
	utils.NewLine(1)
	fmt.Println("1. Check Balance \t 2. Withdraw")
	fmt.Println("3. Deposit \t\t 4. Change PIN")
	fmt.Println("0. Exit")

	var selectedOption int
	_, err := fmt.Scan(&selectedOption)
	utils.CheckErr(err)

	switch selectedOption {
	case 1:
		CheckBalance()
	case 2:
		Withdraw()
	case 3:
		Deposit()
	case 4:
		ChangePIN()
	case 0:
		Exit()
	default:
		fmt.Println("ERROR: invalid input, select a valid operation")
		utils.NewLine(1)
		Menu()
	}

}

// ChackBalance displays current user's account balance.
func CheckBalance() {
	utils.NewLine(1)
	fmt.Println("----------WITHDRAWABLE BALANCE----------")
	fmt.Printf("\t\t ₦%.2f", acctBalance)
	utils.NewLine(2)
	NewOperation()
}

// Deposit handle cash deposit operations for user.
func Deposit() {
	utils.NewLine(1)
	fmt.Println("--------------DEPOSIT----------------")
	utils.NewLine(1)

	var depositAmount float64
	var selectedOption int

	fmt.Print(" How much would you like to deposit?")
	utils.NewLine(1)
	fmt.Println("1. 5000 \t\t 2. 10000")
	fmt.Println("3. 20000 \t\t 4. 30000")
	fmt.Println("5. 50000\t\t 6. other")

	_, err := fmt.Scan(&selectedOption)
	utils.NewLine(1)
	utils.CheckErr(err)

	for {
		switch selectedOption {
		case 1:
			acctBalance += 5000
			fmt.Println("You have successfully deposited ₦5000")
		case 2:
			acctBalance += 10000
			fmt.Println("You have successfully deposited ₦10000")
		case 3:
			acctBalance += 20000
			fmt.Println("You have successfully deposited ₦20000")
		case 4:
			acctBalance += 30000
			fmt.Println("You have successfully deposited ₦30000")
		case 5:
			acctBalance += 50000
			fmt.Println("You have successfully deposited ₦50000")
		case 6:
			fmt.Print("Enter the amount you wish to deposit: ")
			for {
				_, err := fmt.Scan(&depositAmount)
				utils.CheckErr(err)
				if depositAmount > 0 {
					acctBalance += depositAmount
					fmt.Printf("You have successfully deposited ₦%.2f", depositAmount)
					utils.NewLine(1)
					break
				} else {
					fmt.Printf("ERROR: cannot deposit ₦%.2f, enter valid amount: ", depositAmount)
				}
			}

		default:
			fmt.Println("ERROR: invalid input, select a valid option")
		}

		NewOperation()

	}

}

// Withdraw handles cash withdrawal operations for user.
func Withdraw() {
	utils.NewLine(1)
	fmt.Println("--------------WITHDRAWAL----------------")
	utils.NewLine(1)

	var withdrawalAmount float64
	var selectedOption int

	fmt.Print(" How much would you like to withdraw?")
	utils.NewLine(1)
	fmt.Println("1. 5000 \t\t 2. 10000")
	fmt.Println("3. 15000 \t\t 4. 20000")
	fmt.Println("5. 30000\t\t 6. other")

	_, err := fmt.Scan(&selectedOption)
	utils.NewLine(1)
	utils.CheckErr(err)

	for {
		if acctBalance == 0 {
			fmt.Println("ERROR: account balance is currently ₦0.00, please make a deposit")
			Menu()
		}

		if withdrawalAmount > acctBalance {
			fmt.Println("Insufficient Funds")
			NewOperation()
		}

		switch selectedOption {
		case 1:
			acctBalance -= 5000
			fmt.Println("You have successfully withdrawn ₦5000")
		case 2:
			acctBalance -= 10000
			fmt.Println("You have successfully withdrawn ₦10000")
		case 3:
			acctBalance -= 15000
			fmt.Println("You have successfully withdrawn ₦20000")
		case 4:
			acctBalance -= 20000
			fmt.Println("You have successfully withdrawn ₦30000")
		case 5:
			acctBalance -= 30000
			fmt.Println("You have successfully withdrawn ₦50000")
		case 6:
			fmt.Print("Enter the amount you wish to withdraw: ")
			for {
				_, err := fmt.Scan(&withdrawalAmount)
				utils.CheckErr(err)
				if withdrawalAmount > 0 {
					if acctBalance == 0 {
						fmt.Println("ERROR: account balance is currently ₦0.00, please make a deposit")
						Menu()
					}

					if withdrawalAmount > acctBalance {
						fmt.Println("Insufficient Funds")
						NewOperation()
					}

					acctBalance -= withdrawalAmount
					fmt.Printf("You have successfully withdrawn ₦%.2f", withdrawalAmount)
					utils.NewLine(1)
					break
				} else {
					fmt.Printf("ERROR: cannot withdraw ₦%.2f, enter valid amount: ", withdrawalAmount)
				}
			}

		default:
			fmt.Println("ERROR: invalid input, select a valid option")
		}

		NewOperation()

	}
}

// NewOperation prompts user to either perform a
// new operation or end the session.
func NewOperation() {
	var input string
	var parsedInput string

	fmt.Print("Would you like to perform anther transaction? (y/n): ")
	fmt.Scan(&input)
	parsedInput = strings.ToLower(strings.TrimSpace(input))

	switch parsedInput {
	case "y":
		Menu()
	case "n":
		Exit()
	default:
		fmt.Println("ERROR: invalid input, choose a valid option")
		utils.NewLine(1)
		NewOperation()
	}
}

// Exit terminates running user session.
func Exit() {
	utils.NewLine(1)
	fmt.Println("Thank you for banking with us.")
	os.Exit(0)
}
