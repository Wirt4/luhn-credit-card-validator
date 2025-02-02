package server

import (
	"fmt"
	"log"
	"net/http"

	"main.go/packages/handlers"
	"main.go/packages/luhn"
	"main.go/packages/routing_checksum"
)

type Server struct {
	port string
}

func NewServer(port string) *Server {
	return &Server{port: port}
}

func (s *Server) ListenAndServe() {
	credit_card_handler := handlers.NewHandler(&luhn.LuhnValidator{})
	routing_number_handler := handlers.NewHandler(&routing_checksum.RoutingChecksum{})
	http.HandleFunc(validate("credit_card"), credit_card_handler.HandleRequest)
	http.HandleFunc(validate("routing_number"), routing_number_handler.HandleRequest)
	fmt.Println("Server is listening on port", s.port)
	log.Fatal(http.ListenAndServe(":"+s.port, nil))
}

func validate(s string) string {
	a := "/validate/" + s + "/"
	return a
}
