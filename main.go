package main

import (
	"log"

	"github.com/aristogram/user/internal/server"
)

func main() {
	s := server.NewServer(":14888")
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
