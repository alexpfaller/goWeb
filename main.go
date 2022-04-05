package main

import (
	"net/http"

	"github.com/gookit/color"
)

var userInput byte
var menuOutcome byte
var runningPort int32

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
		startServer()
	}
}

// menu()
// menuValidation()

func startServer() {
	http.Handle("/", http.FileServer(http.Dir("default")))
	http.ListenAndServe(":8090", nil)
}
