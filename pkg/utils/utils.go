package utils

import (
	"fmt"
	"log"
)

// NewLine prints number of new lines based on argument passed.
func NewLine(n int) {
	for ; n > 0; n-- {
		fmt.Println()
	}
}

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
