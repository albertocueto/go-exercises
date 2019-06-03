package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "password12345"
	ep, error := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if error != nil {
		fmt.Println("There was an error with bcrypt:", error)
	}
	fmt.Println(password, ep)
	// passwordToCheck := "PaSsWoRd123"
	passwordToCheck := "password12345"
	error = bcrypt.CompareHashAndPassword(ep, []byte(passwordToCheck))
	if error != nil {
		fmt.Println("Passwords don't match")
		return
	}
	fmt.Println("Welcome")
}
