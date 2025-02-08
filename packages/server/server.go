package server

import (
	"fmt"
	"log"
	"net/http"

	"main.go/packages/factories"
	"main.go/packages/handlers"
	"main.go/packages/luhn"
)

type Server struct {
	port string
}

func NewServer(port string) *Server {
	return &Server{port: port}
}

func (s *Server) ListenAndServe() {
	handler := handlers.NewHandler(&luhn.LuhnValidator{}, &factories.CreditCardFactory{})
	http.HandleFunc("/validate_credit_card/", handler.HandleRequest)
	fmt.Println("Server is listening on port", s.port)
	log.Fatal(http.ListenAndServe(":"+s.port, nil))
}
