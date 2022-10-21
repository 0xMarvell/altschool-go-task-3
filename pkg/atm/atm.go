package atm

import (
	"fmt"
	"os"
)

// WelcomeScreen displays welcome message.
func WelcomeScreen() {

}

// Menu displays all available operations to user.
func Menu() {

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
