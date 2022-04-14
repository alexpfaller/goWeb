package main

import (
	"fmt"

	"github.com/gookit/color"
)

func Menu() {
	yellow := color.FgYellow.Render
	fmt.Printf("%sTo configure webserver\n", yellow("[1]"))
	fmt.Printf("%sTo start webserver\n", yellow("[2]"))
	fmt.Printf("%sAbort\n", yellow("[3]"))
	fmt.Scan(&userInput)
}
func MenuValidation() byte {
	if userInput < 1 || userInput > 3 {
		color.Error.Println("Invalid input!")
		color.Error.Println("Please chose one of the above given options..")
		main()
	} else if userInput == 1 {
		menuOutcome = 1
	} else if userInput == 2 {
		menuOutcome = 2
	} else if userInput == 3 {
		menuOutcome = 3
	}
	return menuOutcome
}
