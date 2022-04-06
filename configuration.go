package main

import (
	"fmt"

	"github.com/gookit/color"
)

var SubPageInput int
var Subpages = make([]int, SubPageInput)

func ConfigServer() {
	color.Success.Println("Webserver configuration menu")
	for {
		ConfigMenu()
		ConfigmenuValidation()
		if menuOutcome >= 1 || menuOutcome <= 3 {
			break
		}
	}
	switch menuOutcome {
	case 1:
		SetSubPage()
	case 2:
		SetPort()
	case 3:
		fmt.Println("To create redirect")
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
func SetSubPage() {
	fmt.Printf("Enter amount of subpages you want to create: ")
	fmt.Scan(&SubPageInput)
	for i := 0; i < SubPageInput; i++ {
		fmt.Printf("Enter %dth element: ", i)
		fmt.Scanf("%d", &Subpages[i])
	}
}
