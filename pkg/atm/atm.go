package atm

import (
	"fmt"
	"os"
)

// WelcomeScreen displays welcome message.
func WelcomeScreen() {
	fmt.Println("------------------------")
	fmt.Println("          ATM          ")
	fmt.Println("------------------------")

}

// Menu displays all available operations to user.
func Menu() {
	fmt.Println("AVAILABLE OPERATIONS")
	fmt.Println("1. Check Balance \t 2. Withdraw")
	fmt.Println("3. Deposit \t\t 4. Change PIN")
	fmt.Println("\t\t\t 0. Exit")

}

// ChackBalance displays current user's account balance.
func CheckBalance() {

}

// Deposit handle cash deposit operations for user.
func Deposit() {

}

// Withdraw handles cash withdrawal operations for user.
func Withdraw() {

}

// Exit terminates the running user session.
func Exit() {
	fmt.Println("Thank you for banking with us.")
	os.Exit(0)
}
