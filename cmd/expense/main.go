package main

import (
	"fmt"
	"github.com/expense-tracker-cli/internal/password"
	"github.com/fatih/color"
	"os"
)

func main() {
	errclr := color.New(color.BgRed).Add(color.FgWhite)
	var pwd string
	fmt.Println("Enter the password: ")
	fmt.Scanln(&pwd)
	pwd_crct := password.Authenticate(pwd)
	if pwd_crct {
		main_program()
	} else {
		errclr.Println("Wrong Password!")
		os.Exit(1)
	}
}
