package main

import (
	"main.go/packages/server"
)

func main() {
	s := server.NewServer("8080")
	s.ListenAndServe()

}
