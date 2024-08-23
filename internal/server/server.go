package server

import (
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

func NewServer(address string) *Server {
	s := &Server{
		mux:     http.NewServeMux(),
		address: address,
	}

	for address, handler := range Endpoints {
		s.mux.Handle(address, handler)
	}

	return s
}

func (s *Server) Start() error {
	return http.ListenAndServe(s.address, s.mux)
}
