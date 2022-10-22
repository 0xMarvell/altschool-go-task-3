package atm

import (
	"fmt"

	"github.com/0xMarvell/altschool-go-task-3/pkg/utils"
)

var PIN string = "0000"

// Login handles user login.
func Login() {
	for {
		if !VerifyPIN(PIN) {
			continue
		}

		break
	}
}

// ChangePIN updates userr account PIN.
func ChangePIN() {
	utils.NewLine(1)
	fmt.Println("change pin")
	for {
		fmt.Printf("Enter new pin: ")
		var updatedPIN string
		fmt.Scan(&updatedPIN)

		if len(updatedPIN) != 4 {
			fmt.Println("Pin should be 4 characters long")
			continue
		}

		PIN = updatedPIN
		fmt.Println("Pin has been changed")
		break
	}
}

// VerifyPIN checks if inputted PIN matches user account PIN
func VerifyPIN(acctPIN string) bool {
	var inputPIN string
	fmt.Printf("ENTER PIN: ")
	fmt.Scan(&inputPIN)

	if acctPIN != inputPIN {
		fmt.Println("Incorrect pin")
		return false
	}

	return true
}
