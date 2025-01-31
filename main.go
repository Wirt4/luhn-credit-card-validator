package main

import (
	"fmt"
	"log"
	"net/http"

	"main.go/packages/handlers"
	"main.go/packages/luhn"
)

func main() {
	handler := handlers.NewHandler(&luhn.LuhnValidator{})
	http.HandleFunc("/validate_credit_card/", handler.HandleRequest)
	fmt.Println("Server is listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
