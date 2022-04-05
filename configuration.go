package main

import (
	"fmt"

	"github.com/gookit/color"
)

func ConfigServer() {
	color.Success.Println("Webserver configuration menu")
	for {
		ConfigMenu()
		ConfigmenuValidation()
		if menuOutcome >= 1 || menuOutcome <= 3 {
			break
		}
	}
}
func ConfigMenu() {
	yellow := color.FgYellow.Render
	fmt.Printf("%sTo set sub-page\n", yellow("[1]"))
	fmt.Printf("%sTo set port\n", yellow("[2]"))
	fmt.Printf("%sTo create redirect\n", yellow("[3]"))
	fmt.Scan(&userInput)
}
func ConfigmenuValidation() byte {
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
