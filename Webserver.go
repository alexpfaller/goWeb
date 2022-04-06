package main

import "net/http"

func StartServer(ServerPort string) {
	http.Handle("/", http.FileServer(http.Dir("default")))
	http.ListenAndServe(ServerPort, nil)
}
