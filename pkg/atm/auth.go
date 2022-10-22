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
	fmt.Println("-------------CHANGE PIN-------------")
	for {
		fmt.Printf("Enter new PIN: ")
		var updatedPIN string
		_, err := fmt.Scan(&updatedPIN)
		utils.CheckErr(err)

		if len(updatedPIN) != 4 {
			fmt.Println("PIN should be 4 characters long")
			continue
		}

		PIN = updatedPIN
		fmt.Println("PIN updated successfully")
		NewOperation()
		break
	}
}

// VerifyPIN checks if inputted PIN matches user account PIN
func VerifyPIN(acctPIN string) bool {
	var inputPIN string
	fmt.Printf("ENTER PIN: ")
	_, err := fmt.Scan(&inputPIN)
	utils.CheckErr(err)

	if acctPIN != inputPIN {
		fmt.Println("Incorrect PIN")
		return false
	}

	return true
}
