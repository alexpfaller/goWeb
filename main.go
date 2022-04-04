package main

import (
	"fmt"
	"net/http"

	"github.com/gookit/color"
)

var userInput byte
var menuOutcome byte
var runningPort int32

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello workd")
}

func main() {
	color.Cyan.Printf("Welcome to goWeb\n\n\n")
	for {
		menu()
		menuValidation()
		if menuOutcome >= 1 || menuOutcome <= 3 {
			break
		}
	}
	switch menuOutcome {
	case 1:
		configServer()
	case 2:
		color.Success.Println("Your server is now ap and running...")
		startServer()
	}
}

func menu() {
	yellow := color.FgYellow.Render
	fmt.Printf("%sTo configure webserver\n", yellow("[1]"))
	fmt.Printf("%sTo start webserver\n", yellow("[2]"))
	fmt.Printf("%sAbort\n", yellow("[3]"))
	fmt.Scan(&userInput)
}
func menuValidation() byte {
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
func configServer() {
	color.Success.Println("Webserver configuration menu")
	for {
		configMenu()
		configmenuValidation()
		if menuOutcome >= 1 || menuOutcome <= 3 {
			break
		}
	}
}
func configMenu() {
	yellow := color.FgYellow.Render
	fmt.Printf("%sTo set sub-page\n", yellow("[1]"))
	fmt.Printf("%sTo set port\n", yellow("[2]"))
	fmt.Printf("%sTo create redirect\n", yellow("[3]"))
	fmt.Scan(&userInput)
}
func configmenuValidation() byte {
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
func startServer() {
	http.Handle("/", http.FileServer(http.Dir("default")))
	http.ListenAndServe(":8090", nil)
}
