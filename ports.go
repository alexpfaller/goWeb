package main

import (
	"fmt"

	"github.com/gookit/color"
)

func SetPort() {
	color.Success.Println("Port configuration menu")
	for {
		PortMenu()
		PortmenuValidation()
		if menuOutcome >= 1 || menuOutcome <= 3 {
			break
		}
	}
}
func PortMenu() {
	yellow := color.FgYellow.Render
	fmt.Printf("%sTo scan for open ports\n", yellow("[1]"))
	fmt.Printf("%sTo set auto port\n", yellow("[2]"))
	fmt.Printf("%sTo set static port\n", yellow("[3]"))
	fmt.Scan(&userInput)
}
func PortmenuValidation() byte {
	if userInput < 1 || userInput > 3 {
		color.Error.Println("Invalid input!")
		color.Error.Println("Please chose one of the above given options..")
	} else if userInput == 1 {
		menuOutcome = 1
	} else if userInput == 2 {
		menuOutcome = 2
	} else if userInput == 3 {
		menuOutcome = 3
	}
	return menuOutcome
}
func DefaultPort() {

}
