package main

import (
	"github.com/gookit/color"
)

var userInput uint16
var menuOutcome byte

func main() {
	color.Cyan.Printf("Welcome to goWeb\n\n\n")
	for {
		Menu()
		MenuValidation()
		if menuOutcome >= 1 || menuOutcome <= 3 {
			break
		}
	}
	switch menuOutcome {
	case 1:
		ConfigServer()
	case 2:
		color.Success.Println("Your server is now ap and running...")
		StartServer(ServerPort)
	}
}

// menu()
// menuValidation()
