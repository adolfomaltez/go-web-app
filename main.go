package main

import (
	"fmt"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello OneEuronet")
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
}

func main() {
	fmt.Println("GO Web App Started on Port 3000")
	setupRoutes()
	http.ListenAndServe(":3000", nil)
}
