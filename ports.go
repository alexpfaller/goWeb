package main

import (
	"fmt"

	"github.com/gookit/color"
)

var defaultPort uint16 = 8090
var ServerPort string
var ServerPortZwei string

func SetPort() {
	color.Success.Println("Port configuration menu")
	for {
		PortMenu()
		PortmenuValidation()
		if menuOutcome >= 1 || menuOutcome <= 3 {
			break
		}
	}
	switch menuOutcome {
	case 1:
		fmt.Println("to scan for open ports")
	case 2:
		fmt.Println("to set auto port")
	case 3:
		SetStaticPort()
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
		SetPort()
	} else if userInput == 1 {
		menuOutcome = 1
	} else if userInput == 2 {
		menuOutcome = 2
	} else if userInput == 3 {
		menuOutcome = 3
	}
	return menuOutcome
}
func SetStaticPort() {
	yellow := color.FgYellow.Render
	fmt.Printf("%sSet your own static port\n", yellow("[1]"))
	fmt.Printf("%sUse default static port (:%v)\n", yellow("[2]"), defaultPort)
	fmt.Scan(&userInput)
	switch userInput {
	case 1:
		CustomStaticPort()
	case 2:
		DefaultPort()
	}
}
func CustomStaticPort() {
	fmt.Scan(&userInput)
	ServerPort = fmt.Sprintf(":%v", userInput)
	color.Success.Println("Your server is now up and running...")
	color.Green.Printf("http://127.0.0.1%v", ServerPort)
	StartServer(ServerPort)
}
func DefaultPort() {
	ServerPort = ":8090"
	color.Success.Println("Your server is now up and running...")
	color.Green.Printf("http://127.0.0.1%v", ServerPort)
	StartServer(ServerPort)
}
