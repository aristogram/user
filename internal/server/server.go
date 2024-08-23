package server

import (
	"log"
	"net/http"
)

var Endpoints = map[string]http.Handler{
	"/ping": http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "ты лох", 200)
	}),
}

type Server struct {
	mux     *http.ServeMux
	address string
}

func NewServer() *Server {
	s := &Server{
		mux: http.NewServeMux(),
	}

	for address, handler := range Endpoints {
		s.mux.Handle(address, handler)
	}

	return s
}

func (s *Server) Start(address string) {
	s.address = address

	err := http.ListenAndServe(s.address, s.mux)
	if err != nil {
		log.Fatalln("server error:", err)
	}
}
