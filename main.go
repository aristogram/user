package main

import "github.com/aristogram/user/internal/server"

func main() {
	s := server.NewServer()
	s.Start(":14888")
}
