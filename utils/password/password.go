// Package password provides Pass_checker, Password_compare
package password

import (
	"fmt"
	"os"
)

func Password_compare(pass string) bool {
	realpass := os.Getenv("TGBOTPASS")
	fmt.Println(realpass)
	return (realpass == pass)
}

// password correction checking
func Pass_checker(result *string, pass string) bool {
	if !Password_compare(pass) {
		*result = "password wrong"
		return false
	} else {
		*result = "correct! you gain admin access"
		return true
	}
}
