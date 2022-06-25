package pass

import (
	"fmt"
	"os"
	"strings"
	"syscall"

	"golang.org/x/term"
)

func GetPassword() string {
	for {
		fmt.Print("Enter password: ")
		bytepw, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			os.Exit(1)
		}
		pass1 := string(bytepw)
		fmt.Print("\n")

		fmt.Print("Repeat password: ")
		bytepw, err = term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			os.Exit(1)
		}
		pass2 := string(bytepw)
		fmt.Print("\n")

		if strings.Compare(pass1, pass2) == 0 {
			return pass1
		} else {
			fmt.Println("The passwords entered do not match. Please try again.")
		}
	}
}
